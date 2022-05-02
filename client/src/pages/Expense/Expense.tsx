import React from "react";

const Expense = () => {
  return (
    <div>
      <h2>Expense</h2>
      <form action="">
        <input type="text" placeholder="person" />
        <br />
        <br />
        <input type="text" placeholder="productname" />
        <br />
        <br />
        <input type="text" placeholder="purchase" />
        <button type="submit">submit</button>
      </form>
      <p>person</p>
      <p>productname</p>
      <p>purchase</p>
    </div>
  );
};

export default Expense;
