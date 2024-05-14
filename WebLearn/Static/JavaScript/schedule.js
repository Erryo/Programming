function SubForm() {
  $.ajax({
    url: "/Submit/AddSubject",
    type: "POST",
    data: $("#myForm").serialize(),
    success: function () {
      alert("aa");
    },
  });
  console.log(document.getElementById("SubjField").value);
}
