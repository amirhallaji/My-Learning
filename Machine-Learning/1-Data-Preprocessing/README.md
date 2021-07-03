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
