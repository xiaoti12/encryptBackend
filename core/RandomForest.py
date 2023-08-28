import joblib
# from sklearn.preprocessing import MinMaxScaler

import FeatureConstant


class RandomForestPredictor():
    def __init__(self, model_path, scaler_path) -> None:
        with open(model_path, "rb") as f:
            self.model = joblib.load(f)
        with open(scaler_path, "rb") as f:
            self.scaler = joblib.load(f)

    def predict(self, df):
        df = self.preprocess(df)
        result = self.model.predict_proba(df)
        return result

    def preprocess(self, df):
        df = df[FeatureConstant.FEATURE].values
        df_scaler = self.data_tranform(df)
        return df_scaler

    def data_tranform(self, df):
        return self.scaler.transform(df)


if __name__ == "__main__":
    import os
    dir_path=os.path.dirname(os.path.abspath(__file__))
    model_path = os.path.join(dir_path,"train","rfc_nomkv.joblib")
    scaler_path = os.path.join(dir_path,"train","scaler_nomkv.joblib")
    rfc = RandomForestPredictor(model_path, scaler_path)

    import parseTLS
    data = parseTLS.log_to_df(logName='/home/fsc/liujy/realtime/tls.log')
    result = rfc.predict(data)
    print(result)
