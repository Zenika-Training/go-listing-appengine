(function () {
  const items = document.querySelectorAll("#list li[data-name]");
  const input = document.getElementById("q");
  const count = document.getElementById("count");
  const none = document.getElementById("none");
  const norm = (s) => s.toLowerCase().normalize("NFD").replace(/[̀-ͯ]/g, "");
  function filter() {
    const q = norm(input.value.trim());
    let n = 0;
    items.forEach((li) => {
      const ok = norm(li.dataset.name).indexOf(q) !== -1;
      li.hidden = !ok;
      if (ok) n++;
    });
    count.textContent = items.length ? n + " formation" + (n > 1 ? "s" : "") + (q ? " trouvée" + (n > 1 ? "s" : "") : " disponible" + (n > 1 ? "s" : "")) : "";
    none.hidden = n !== 0 || !items.length;
  }
  input.addEventListener("input", filter);
  filter();
})();
