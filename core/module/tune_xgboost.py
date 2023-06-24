# -*- coding: utf-8 -*-

import xgboost
from sklearn.model_selection import GridSearchCV
from xgboost import XGBClassifier
from sklearn import metrics
from module.model_judge import model_judge
from sklearn.metrics import silhouette_score as sc

def cv_silhouette_scorer(estimator, X):
    estimator.fit(X)
    cluster_labels = estimator.labels_
    num_labels = len(set(cluster_labels))
    num_samples = len(X.index)
    if num_labels == 1 or num_labels == num_samples:
        return -1
    else:
        return sc(X, cluster_labels)

cv = [(slice(None), slice(None))]

# 使用交叉验证调整n_estimators
def modelfit(alg, dtrain, dtest, logger, name, cv_folds=2, early_stopping_rounds=50):
    """
    使用交叉验证调整n_estimators

    参数
    ----------
    alg: XGBClassifier模型
    dtrain: 训练数据
    dtest: 测试数据
    cv_fold: 交叉验证折数(default:5)
    early_stopping_rounds:(default:50)

    返回
    ----------
    n_estimators
    """
    xgb_param = alg.get_xgb_params()
    xgtrain = xgboost.DMatrix(dtrain[:, :302], label=dtrain[:, 302:])
    cvresult = xgboost.cv(xgb_param, xgtrain, num_boost_round=alg.get_params()['n_estimators'], nfold=cv_folds,
                           metrics='mlogloss', early_stopping_rounds=early_stopping_rounds, seed=27, shuffle=False)
    alg.set_params(n_estimators=cvresult.shape[0])
    logger.info(f"Select n_estimators = {cvresult.shape[0]}")

    # Fit the algorithm on the data
    alg.fit(dtrain[:, :302], dtrain[:, 302:], eval_metric='mlogloss')

    # Predict training set:
    dtrain_predictions = alg.predict(dtrain[:, :302])
    dtest_predictions = alg.predict(dtest[:, :302])

    logger.info(f"\nModel Report")
    logger.info("Accuracy(Train) : %.4g" %
                metrics.accuracy_score(dtrain[:, 302:], dtrain_predictions))
    logger.info("Accuracy(Test) : %.4g" %
                metrics.accuracy_score(dtest[:, 302:], dtest_predictions))

    logger.info(f"在训练集中")
    model_judge(dtrain[:, 302:], dtrain_predictions, logger, name + '_train')
    logger.info(f"在验证集中")
    model_judge(dtest[:, 302:], dtest_predictions, logger, name + '_test')

    return cvresult.shape[0]


# 输入: data_train:训练数据 data_test:验证数据
def tune_xgboost(data_train, data_test, logger):
    """
    训练XGBoost模型并进行调参

    参数
    ----------
    data_train:训练数据(含标签)
    data_test:测试数据(含标签)

    返回:
    ----------
    final_xgb:调参完毕的XGBClassifier模型
    """
    # Step 1: Fix learning rate and number of estimators for tuning tree-based parameters
    logger.info(
        f"-----Step 1: Fix learning rate and number of estimators for tuning tree-based parameters-----")
    params = {"learning_rate":0.1,"n_estimators":1000,"max_depth":5,"min_child_weight":1,"gamma":0,"subsample":0.8,
              "colsample_bytree":0.8,"n_jobs":-1,"seed":27,"objective":"multi:softmax","num_class":6,
              "gpu_id":0, "tree_method":"gpu_hist"}
    xgb1 = XGBClassifier(**params)
    xgb1_n_estimators = modelfit(xgb1, data_train, data_test,logger,'xgb1')
    logger.info(f"Update n_estimators to {xgb1_n_estimators}")
    params["n_estimators"] = xgb1_n_estimators

    # Step 2: Tune max_depth and min_child_weight
    logger.info(f"-----Step 2: Tune max_depth and min_child_weight-----")
    param_test1 = {
        'max_depth': range(3, 10, 2),
        'min_child_weight': range(1, 6, 2)
    }
    
    gsearch1 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test1, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch1.fit(data_train[:, :302], data_train[:, 302:])
    # gsearch1.cv_results_
    logger.info(f"The best params is {gsearch1.best_params_}")
    logger.info(f"The best f1 score is {gsearch1.best_score_}")

    g1_max_depth = gsearch1.best_params_['max_depth']
    g1_min_child_weight = gsearch1.best_params_['min_child_weight']
    param_test2 = {
        'max_depth': ([1, 2, 3] if g1_max_depth == 1 else [g1_max_depth - 1, g1_max_depth, g1_max_depth + 1]),
        'min_child_weight': ([1, 2, 3] if g1_min_child_weight == 1 else [g1_min_child_weight - 1, g1_min_child_weight,
                                                                         g1_min_child_weight + 1])
    }
    gsearch2 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test2, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch2.fit(data_train[:, :302], data_train[:, 302:])
    logger.info(f"The best params is {gsearch2.best_params_}")
    logger.info(f"The best f1 score is {gsearch2.best_score_}")
    params["max_depth"] = gsearch2.best_params_['max_depth']
    params["min_child_weight"] = gsearch2.best_params_['min_child_weight']

    # Step 3: Tune gamma
    logger.info(f"-----Step 3: Tune gamma-----")
    param_test3 = {
        'gamma': [i / 10.0 for i in range(0, 5)]
    }
    gsearch3 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test3, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch3.fit(data_train[:, :302], data_train[:, 302:])
    logger.info(f"The best params is {gsearch3.best_params_}")
    logger.info(f"The best f1 score is {gsearch3.best_score_}")
    params["gamma"] = gsearch3.best_params_['gamma']

    # reset n_estimators
    params["n_estimators"] = 1000
    xgb2 = XGBClassifier(**params)
    xgb2_n_estimators = modelfit(xgb2, data_train, data_test,logger,'xgb2')
    logger.info(f"Update n_estimators to {xgb2_n_estimators}")
    params["n_estimators"] = xgb2_n_estimators

    # Step 4: Tune subsample and colsample_bytree
    logger.info(f"-----Step 4: Tune subsample and colsample_bytree-----")
    param_test4 = {
        'subsample': [i / 10.0 for i in range(6, 10)],
        'colsample_bytree': [i / 10.0 for i in range(6, 10)]
    }
    gsearch4 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test4, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch4.fit(data_train[:, :302], data_train[:, 302:])
    logger.info(f"The best params is {gsearch4.best_params_}")
    logger.info(f"The best f1 score is {gsearch4.best_score_}")

    g4_colsample_bytree = gsearch4.best_params_['colsample_bytree']
    g4_subsample = gsearch4.best_params_['subsample']
    param_test5 = {
        'subsample': [i / 100.0 for i in
                      range((0 if g4_subsample - 0.1 <= 0 else int((g4_subsample - 0.1) * 100)),
                            (100 if g4_subsample + 0.1 > 100 else int((g4_subsample + 0.1) * 100)), 5)],
        'colsample_bytree': [i / 100.0 for i in
                             range((0 if g4_colsample_bytree - 0.1 <= 0 else int((g4_colsample_bytree - 0.1) * 100)),
                                   (100 if g4_colsample_bytree + 0.1 >
                                    100 else int((g4_colsample_bytree + 0.1) * 100)),
                                   5)]
    }
    gsearch5 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test5, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch5.fit(data_train[:, :302], data_train[:, 302:])
    logger.info(f"The best params is {gsearch5.best_params_}")
    logger.info(f"The best f1 score is {gsearch5.best_score_}")
    params["subsample"] = gsearch5.best_params_['subsample']
    params["colsample_bytree"] = gsearch5.best_params_['colsample_bytree']

    # Step 5: Tuning Regularization Parameters
    logger.info("-----Step 5: Tuning Regularization Parameters-----")
    param_test6 = {
        'reg_alpha': [0, 0.001, 0.005, 0.01, 0.05, 0.1, 0.2, 1, 100]
    }
    gsearch6 = GridSearchCV(estimator=XGBClassifier(**params),
                            param_grid=param_test6, scoring='accuracy', n_jobs=5, cv=cv)
    gsearch6.fit(data_train[:, :302], data_train[:, 302:])
    logger.info(f"The best params is {gsearch6.best_params_}")
    logger.info(f"The best f1 score is {gsearch6.best_score_}")
    params["reg_alpha"] = gsearch6.best_params_['reg_alpha']

    # Update n_estimators
    logger.info("-----Update n_estimators-----")
    params["n_estimators"] = 1000
    xgb3 = XGBClassifier(**params)
    xgb3_n_estimators = modelfit(xgb3, data_train, data_test, logger, 'xgb3')

    logger.info(f"Update n_estimators to {xgb3_n_estimators}")
    params["n_estimators"] = xgb3_n_estimators

    # Step 6: Reducing Learning Rate
    logger.info(f"-----Step 6: Reducing Learning Rate-----")
    params["n_estimators"] = 5000
    params["learning_rate"] = 0.01
    xgb4 = XGBClassifier(**params)
    xgb4_n_estimators = modelfit(xgb4, data_train, data_test, logger, 'xgb4')
    logger.info(
        f"Update n_estimators in learning_rate=0.01 to {xgb4_n_estimators}")

    # Step 7: Feature selection by RFECV
    params["n_estimators"] = xgb4_n_estimators
    final_xgb = XGBClassifier(**params)

    logger.info(
        f"Final singleXGBoost parameters is learning_rate=0.01, n_estimators={xgb4_n_estimators}, max_depth={gsearch2.best_params_['max_depth']}, min_child_weight={gsearch2.best_params_['min_child_weight']}, gamma={gsearch3.best_params_['gamma']}, subsample={gsearch5.best_params_['subsample']}, colsample_bytree={gsearch5.best_params_['colsample_bytree']}, reg_alpha={gsearch6.best_params_['reg_alpha']}")
    return final_xgb
