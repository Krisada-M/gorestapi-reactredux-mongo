import { createSlice } from "@reduxjs/toolkit";

const incomes = createSlice({
  name: "incomes",
  initialState: [
    {
      id: "",
      category: "",
      person: "",
      productname: "",
      purchase: 0,
      date: {
        day: "",
        month: "",
        year: "",
        time: "",
      },
    },
  ],
  reducers: {
    getIncomeSlice: (state, action) => {
      state = action.payload;
      return state;
    },
    addIncomeSlice: (state, action) => {
      state.push(action.payload);
      return state;
    },
    editIncomeSlice: (state, action) => {
      state = state.map((i) =>
        i.id === action.payload.id ? action.payload : i
      );
      return state;
    },
    deleteIncomeSlice: (state, action) => {
      state = state.filter((i) => i.id !== action.payload);
      return state;
    },
  },
});

export const {
  getIncomeSlice,
  addIncomeSlice,
  editIncomeSlice,
  deleteIncomeSlice,
} = incomes.actions;
export default incomes.reducer;
