function SubForm() {
  $.ajax({
    url: "/Submit/AddSubject",
    type: "POST",
    data: $("#myForm").serialize(),
    success: function () {},
  });
  console.log(document.getElementById("SubjField").val());
}
$(document).ready(function () {
  $("#AddSubj").click(function () {
    let val = $("#SubjField").val();
    if (val !== "" || val !== " ") {
      $("#last").before("<li>" + val + "</li>");
    }
  });
});
function DeleteSubj(subject) {
  console.log(subject, "delete");
  $.ajax({
    url: "/Submit/AddSubject?" + $.param({ S: subject }),
    type: "DELETE",
  });
}
function SendSchedule() {
  $.ajax({
    url: "/Submit/AddSchedule",
    type: "POST",
    data: $("#scheduleForm").serialize(),
    success: function () {},
  });
}
