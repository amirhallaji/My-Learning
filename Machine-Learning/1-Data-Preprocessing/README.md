## Data Preprocessing

### Reading the dataset

Most of the datasets consist of two parts.
1. Input Features in the first columns.
2. Dependent Variable Vector in the last column.
   
They have rows and columns and in order to locate these 2 parts mentioned above, **iloc()** will help a lot.

```py
dataframe.iloc[start_row : end_row , start_col : end_col ]
```
for example, if you want to read all the rows and all the last columns except the last one, the following code is the answer.

```py
X = dataframe.iloc[:, :-1]
```
and if you want to read all the rows and the last column, then we have:

```py
y = dataframe.iloc[:, -1]
```
Note: if *.values* is used after iloc, the type of the result is numpy array.


### Removing missing data

#### from scratch

For locating the missing data, **math** library is good, however it can be done by **scikit-learn**.

```py
import math

missing_data_col = dataframe.iloc[:, column_index].values

for i in range(len(missing_data_col)):
      if math.isnan(missing_data_col[i]):
            # do what you want.
```

#### using scikit-learn

Also, we can use scikit-learn to do it.

```py
from sklearn.impute import SimpleImputer
imputer = SimpleImputer(missing_values=np.nan, strategy='mean')
imputer.fit(X[:, 1:3]) # for example column 1 and 2 have missing values.
X[:, 1:3] = imputer.transform(X[:, 1:3])
```

### Encoding data
datas which are not a number, for example country name or colors had better to be encoded before using.

#### Encoding features

we can use scikit-learn to encode input features.
one-hot encoding is a simple way to encode the data. It can be implemented easily. The following code is its implementation using scikit-learn.

```py
from sklearn.compose import ColumnTransformer
from sklearn.preprocessing import OneHotEncoder
col_t = ColumnTransformer(transformers=[('encoder', OneHotEncoder(), [0])], remainder='passthrough')
X = np.array(col_t.fit_transform(X))
```

#### Encoding Labels

```py
from sklearn.preprocessing import LabelEncoder
label = LabelEncoder()
y = label.fit_transform(y)
```

### Splitting dataset to train and test

We have to split dataset to train and test. In most of the times, 80% of the dataset is dedicated to train and the other 20% for test.
Here again, scikit-learn can do it for us the best.

```py
from sklearn.model_selection import train_test_split
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=1)
```

### Feature Scaling

#### Standaradization

$x_{standard}=\frac{x - mean(x)}{standardDeviation(x)}$

#### Normalization

$x_{norm} = \frac{x-min(x)}{x_{max}-x_{min}}$

#### Implementing using scikit-learn

```py
from sklearn.preprocessing import StandardScaler
sc = StandardScaler()
X_train[:, 3:] = sc.fit_transform(X_train[:, 3:])
X_test[:, 3:] = sc.transform(X_test[:, 3:])
```