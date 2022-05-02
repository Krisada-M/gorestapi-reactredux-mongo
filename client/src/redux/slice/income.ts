import { createSlice } from "@reduxjs/toolkit";

const income = createSlice({
  name: "income",
  initialState: {
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

  reducers: {
    setIncomeSlice: (state, action) => {
      state = action.payload;
      return state;
    },
  },
});

export const { setIncomeSlice } = income.actions;
export default income.reducer;
