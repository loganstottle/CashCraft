[...document.querySelectorAll(".buy")].forEach((btn) => {
  btn.onclick = async (e) => {
    const symbol = e.target.id.replace("buy-", "");
    const dollars = parseFloat(prompt("Amount (dollars): "));

    fetch("/buy", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ Symbol: symbol, Dollars: dollars }),
    });

    reload();
  };
});

[...document.querySelectorAll(".sell")].forEach((btn) => {
  btn.onclick = async (e) => {
    const symbol = e.target.id.replace("sell-", "");
    const amount = parseFloat(prompt("Amount (stock units): "));

    fetch("/sell", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ Symbol: symbol, Amount: amount }),
    });

    reload();
  };
});
