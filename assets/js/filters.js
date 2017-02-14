Vue.filter('nl2br', (value) => {
    return value.split("\n").join("<br>")
});
