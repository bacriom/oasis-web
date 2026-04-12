new Swiper(".mySwiper", {
    centeredSlides: true,
    spaceBetween: 20,
    loop: true,

    pagination: {
        el: ".swiper-pagination",
        clickable: true,
    },

    breakpoints: {
        0: {
            slidesPerView: 1.2 // 🔥 móvil: uno principal + peek lateral
        },
        640: {
            slidesPerView: 1.5
        },
        768: {
            slidesPerView: 2 // tablet
        },
        1024: {
            slidesPerView: 3 // desktop
        }
    }
});