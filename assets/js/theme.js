(function () {
  const root = document.documentElement;
  const toggle = document.getElementById("theme");
  const mq = window.matchMedia("(prefers-color-scheme: dark)");
  function isDark() { return root.dataset.theme ? root.dataset.theme === "dark" : mq.matches; }
  function show() { toggle.setAttribute("aria-checked", String(isDark())); }
  toggle.addEventListener("click", () => {
    const next = isDark() ? "light" : "dark";
    // Choosing what the system already gives goes back to following the system.
    const follow = (next === "dark") === mq.matches;
    if (follow) delete root.dataset.theme;
    else root.dataset.theme = next;
    try {
      if (follow) localStorage.removeItem("theme");
      else localStorage.setItem("theme", next);
    } catch (e) {}
    show();
  });
  mq.addEventListener("change", show);
  show();
})();
