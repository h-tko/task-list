Vue.filter('nl2br', (value) => {

    if (value === null || value.length < 1) {
        return value;
    }

    return value.split("\n").join("<br>");
});

Vue.filter('datetime_format', (value) => {
    var format = 'yyyy/mm/dd h:i:s';
    const targetDate = new Date(value);
    format = format.replace(/yyyy/g, targetDate.getFullYear());
    format = format.replace(/mm/g, targetDate.getMonth() + 1);
    format = format.replace(/dd/g, targetDate.getDate());
    format = format.replace(/h/g, targetDate.getHours());
    format = format.replace(/i/g, targetDate.getMinutes());
    format = format.replace(/s/g, targetDate.getSeconds());

    return format;
});
