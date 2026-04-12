function toggleMenu(e) {
    e.stopPropagation(); // evita que el click se propague
    const sidebar = document.getElementById("sidebar");
    sidebar.classList.toggle("active");
}

// Cerrar al hacer click fuera
document.addEventListener("click", function() {
    const sidebar = document.getElementById("sidebar");
    sidebar.classList.remove("active");
});

// Evita que clicks dentro del menú lo cierren
document.getElementById("sidebar").addEventListener("click", function(e) {
    e.stopPropagation();
});