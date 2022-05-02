import { AxiosResponse } from "axios";
import * as api from "../../api";
import { IAddInterface, IncExpModel, IUpdateOnlyInterface } from "../../models";
import * as slice from "../slice/incomes";
import { setIncomeSlice } from "../slice/income";
import { put, takeEvery } from "redux-saga/effects";
import * as type from "../../types";

export function* getIncomeSaga() {
  const incomes: AxiosResponse<IncExpModel> = yield api.getIncome();
  yield put(slice.getIncomeSlice(incomes.data));
}
export function* getIncomeByPersonSaga(action: any) {
  yield api.incomeByPerson(action.person);
  yield put(setIncomeSlice(action.person));
}
export function* createIncomeSaga(action: any) {
  console.log("add!")
  yield api.addIncome(action.income);
  yield put(slice.addIncomeSlice(action.income));
}

export function* updateIncomeSaga(action: any) {
  yield api.updateIncome(action.income);
  yield put(slice.editIncomeSlice(action.income));
}

export function* deleteIncomeByIdSaga(action: any) {
  yield api.deleteIncome(action.id);
  yield put(slice.deleteIncomeSlice(action.id));
}

export function* watchIncomeAsync() {
  yield takeEvery(type.GET_INCOME, getIncomeSaga);
  yield takeEvery(type.GET_INCOME_BY_NAME, getIncomeByPersonSaga);
  yield takeEvery(type.CREATE_INCOME, createIncomeSaga);
  yield takeEvery(type.UPDATE_INCOME_BY_ID, updateIncomeSaga);
  yield takeEvery(type.DELETE_INCOME_BY_ID, deleteIncomeByIdSaga);
}
