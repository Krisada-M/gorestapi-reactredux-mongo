import React from "react";

import type { RootState } from "../../store";
import { setIncomeSlice } from "../../redux/slice/income";

import "./Income.scss";

const Income = () => {
  return (
    <div className="income">
      <h2>Income</h2>
      <input type="text" placeholder="person" />
      <br />
      <br />
      <input type="text" placeholder="productname" />
      <br />
      <br />
      <input type="number" placeholder="purchase" />
      {/* <button onClick={() => handleSubmit()}>submit</button> */}
      {/* {lists.map((item) => (
        <div key={item.id} className="list">
          <p>{item.id}</p>
          <p>{item.category}</p>
          <p>{item.person}</p>
          <p>{item.productname}</p>
          <p>{item.purchase}</p>
          <p>{item.date?.day}</p>
          <p>{item.date?.month}</p>
          <p>{item.date?.year}</p>
          <p>{item.date?.time}</p>
          <button onClick={() => dispatch(setIncomeSlice(item))}>edit</button>
          <button onClick={() => dispatch(deleteIncomeSlice(item.id))}>
            delete
          </button>
        </div>
      ))} */}
    </div>
  );
};

export default Income;
