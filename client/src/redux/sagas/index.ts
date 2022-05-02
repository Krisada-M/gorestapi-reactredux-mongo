import { all } from "redux-saga/effects";
import { watchIncomeAsync } from "./income";

export function* rootSaga() {
    yield all([
        watchIncomeAsync()
    ])
}