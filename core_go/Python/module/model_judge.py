from matplotlib import pyplot as plt
import numpy as np
from sklearn.metrics import  roc_auc_score, roc_curve, confusion_matrix, multilabel_confusion_matrix
from scipy.optimize import brentq
from scipy.interpolate import interp1d
import seaborn as sns
import random

def model_judge(y_true, y_predict, logger, name):
    """
    模型评价

    参数
    ----------
    y_true:真实标签
    y_predict:预测标签
    """
    # con_mat = multilabel_confusion_matrix(y_true, y_predict)

    # for i in range(len(con_mat)):
    #     con_mat_norm = con_mat[i].astype('float') / con_mat[i].sum(axis=1)[:, np.newaxis]     # 归一化
    #     con_mat_norm = con_mat[i].astype('int')

    #     # === plot ===
    #     plt.figure(figsize=(10, 10))
    #     sns.heatmap(con_mat_norm, annot=True, fmt='d', cmap='Blues')

    #     plt.ylim(0, 7)
    #     plt.xlabel('Predicted labels')
    #     plt.ylabel('True labels')
    #     plt.savefig(f"output/multi_mix/{str(name)+'-'+str(i)}.png")
    #     plt.close()

    con_mat = confusion_matrix(y_true, y_predict)
    con_mat_norm = con_mat.astype('int')

    # === plot ===
    plt.figure(figsize=(10, 10))
    sns.heatmap(con_mat_norm, annot=True, fmt='d', cmap='Blues')

    plt.ylim(0, 6)
    plt.xlabel('Predicted labels')
    plt.ylabel('True labels')
    plt.savefig(f"output/multi_mix/{str(name)}.png")
    plt.close()