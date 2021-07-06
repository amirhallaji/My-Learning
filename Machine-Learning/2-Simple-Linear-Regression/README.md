## Simple Linear Regression

### Concept
Simple linear regression is a linear model for predicting data.
As we call it **simple**, it has only one input feature.

It's formula is:

```py
y = b1x + b0
```
**y**: the prediction.
**b1**: the coefficient which is the slope of our linear model.
**b0**: the bias.

### Implementing simple linear regression using scikit-learn

```py
from sklearn.linear_model import LinearRegression
regressor = LinearRegressor()
regressor.fit(X_train, Y_train)
```

for predicting the test result, we have:

```py
y_pred = regressor.predict(X_test)
```