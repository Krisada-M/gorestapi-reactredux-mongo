import { configureStore } from "@reduxjs/toolkit";
import { rootSaga } from "./redux/sagas";
import income from "./redux/slice/income";
import incomes from "./redux/slice/incomes";
import createSagaMiddleware from "@redux-saga/core";
const sagaMiddleware = createSagaMiddleware();

export const store = configureStore({
  reducer: {
    income,
    incomes,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({ thunk: false }).concat(sagaMiddleware),
});
sagaMiddleware.run(rootSaga);
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
