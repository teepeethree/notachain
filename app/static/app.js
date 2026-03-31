// === THEME TOGGLE ===
function toggleTheme() {
  var html = document.documentElement;
  var btn = document.querySelector(".theme-toggle");
  var isDark = html.classList.toggle("dark");
  btn.textContent = isDark ? "Light" : "Dark";
  try {
    localStorage.setItem("notachain-theme", isDark ? "dark" : "light");
  } catch (e) {}
}

// === THEME INIT (respects saved preference, falls back to system) ===
(function () {
  try {
    var saved = localStorage.getItem("notachain-theme");
    var prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    if (saved === "dark" || (!saved && prefersDark)) {
      document.documentElement.classList.add("dark");
      var btn = document.querySelector(".theme-toggle");
      if (btn) btn.textContent = "Light";
    }
  } catch (e) {}
})();

// === FADE-IN OBSERVER ===
var observer = new IntersectionObserver(
  function (entries) {
    entries.forEach(function (entry) {
      if (entry.isIntersecting) {
        entry.target.classList.add("visible");
      }
    });
  },
  { threshold: 0.1 },
);

document.querySelectorAll(".fade-in").forEach(function (el) {
  observer.observe(el);
});
