// Runs before first paint (no defer) so the saved theme is applied without a flash.
document.documentElement.classList.add("js");
try {
  const t = localStorage.getItem("theme");
  if (t === "light" || t === "dark") document.documentElement.dataset.theme = t;
} catch (e) {}
