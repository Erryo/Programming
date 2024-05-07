function ModifyErrorVisibility() {
  const header = document.getElementById("err").innerHTML;
  console.log("Header ");
  if (header != "") {
    document.getElementById("ErrWr").classList.remove("invisible");
    document.getElementById("ErrWr").classList.add("visible");

    console.log("Header is empty");
  } else {
    console.log("Header is not empty");
    document.getElementById("ErrWr").classList.remove("visible");
    document.getElementById("ErrWr").classList.add("invisible");
  }
}
