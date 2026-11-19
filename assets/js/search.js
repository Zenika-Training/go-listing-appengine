(function () {
  var items = document.querySelectorAll("#list li[data-name]");
  var input = document.getElementById("q");
  var count = document.getElementById("count");
  var none = document.getElementById("none");
  function norm(s) { return s.toLowerCase().normalize("NFD").replace(/[̀-ͯ]/g, ""); }
  function filter() {
    var q = norm(input.value.trim()), n = 0;
    items.forEach(function (li) {
      var ok = norm(li.dataset.name).indexOf(q) !== -1;
      li.hidden = !ok;
      if (ok) n++;
    });
    count.textContent = items.length ? n + " formation" + (n > 1 ? "s" : "") + (q ? " trouvée" + (n > 1 ? "s" : "") : " disponible" + (n > 1 ? "s" : "")) : "";
    none.hidden = n !== 0 || !items.length;
  }
  input.addEventListener("input", filter);
  filter();
})();
