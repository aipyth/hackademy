document.querySelectorAll('.dropdown-main').forEach((item) => {
    item.onclick = function (ev) {
        let el = ev.target;
        while (el.localName !== 'div') {
            el = el.parentNode;
        }
        el.nextElementSibling.classList.toggle('hidden');
    };
});
