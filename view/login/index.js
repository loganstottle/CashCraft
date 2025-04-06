document.querySelector("#submit").onclick = async () => {
  const username = document.querySelector("#username").value;
  const password = document.querySelector("#password").value;

  if (!username || !password) {
    return (document.querySelector("#err").innerText =
      "All fields are required.");
  }

  const resp = await fetch("/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });

  try {
    const { error } = await resp.json();
    if (resp.status > 400) document.querySelector("#err").innerText = error;
  } catch {
    location.replace("/");
  }
};
