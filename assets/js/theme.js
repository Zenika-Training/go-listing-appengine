(function () {
  var root = document.documentElement;
  var toggle = document.getElementById("theme");
  var mq = window.matchMedia("(prefers-color-scheme: dark)");
  function isDark() { return root.dataset.theme ? root.dataset.theme === "dark" : mq.matches; }
  function show() { toggle.setAttribute("aria-checked", String(isDark())); }
  toggle.addEventListener("click", function () {
    var next = isDark() ? "light" : "dark";
    // Choosing what the system already gives goes back to following the system.
    var follow = (next === "dark") === mq.matches;
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
