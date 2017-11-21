$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.student-id').append(data.id);
       $('.student-name').append(data.name);
       $('.major').append(data.major);
    });
});
