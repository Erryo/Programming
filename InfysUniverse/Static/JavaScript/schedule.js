function SubForm() {
  $.ajax({
    url: "/Submit/AddSubject",
    type: "POST",
    data: $("#myForm").serialize(),
    success: function () {},
  });
  console.log(document.getElementById("SubjField").value);
}
