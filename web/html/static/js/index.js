function damnPerson(name, gender, level, obscene) {
    const params = $.param({
      name: name,
      gender: gender,
      level: level,
      obscene: obscene,
    });
    window.location.replace(`/damn?${params}`);
}

$(document).ready(function() {

  $("#damnForm").submit(function() {
    const params = $.param({
      name: $("input#name").val(),
      gender: $("select#gender").val(),
      level: $("input#level").val(),
      obscene: $("input#obscene").is(":checked") ? 1 : 0,
    });
    window.location.replace(`/damn?${params}`);

    return false;
  });

});
