let Curr_Player = 0;
let x_score = 0;
let o_score = 0;
let players = ["O", "X"];
function setButton(id) {
  let header = document.getElementById(id);
  if (!players.includes(header.innerHTML)) {
    if (Curr_Player == 0) {
      header.innerHTML = "O";
      Curr_Player = 1;
      header.value = "O";
    } else {
      header.innerHTML = "X";
      header.value = "X";
      Curr_Player = 0;
    }
  }
  if (!checkHoriz()) {
    if (!checkVert()) {
      checkDiago();
    }
  }
  document.getElementById("x_sc").value = x_score;
  document.getElementById("o_sc").value = o_score;
}

function checkHoriz() {
  console.log("H");
  const horiz = [1, 4, 7];
  for (i of horiz) {
    if (!players.includes(document.getElementById(i).value)) {
      continue;
    }
    if (
      document.getElementById(i).value == document.getElementById(i + 1).value
    ) {
      if (
        document.getElementById(i).value == document.getElementById(i + 2).value
      ) {
        document
          .getElementById(i)
          .classList.replace("grid-btn", "grid-btn-red");
        document
          .getElementById(i + 1)
          .classList.replace("grid-btn", "grid-btn-red");
        document
          .getElementById(i + 2)
          .classList.replace("grid-btn", "grid-btn-red");
        disable_all();
        if (document.getElementById(i).value == "X") {
          x_score++;
          return true;
        }
        o_score++;
        return true;
      }
    }
  }
  return false;
}

function checkVert() {
  console.log("V");
  const verti = [1, 2, 3];
  for (i of verti) {
    if (!players.includes(document.getElementById(i).value)) {
      continue;
    }
    if (
      document.getElementById(i).value == document.getElementById(i + 3).value
    ) {
      if (
        document.getElementById(i).value == document.getElementById(i + 6).value
      ) {
        document
          .getElementById(i)
          .classList.replace("grid-btn", "grid-btn-red");
        document
          .getElementById(i + 3)
          .classList.replace("grid-btn", "grid-btn-red");
        document
          .getElementById(i + 6)
          .classList.replace("grid-btn", "grid-btn-red");
        disable_all();
        if (document.getElementById(i).value == "X") {
          x_score++;
          return true;
        }
        o_score++;
        return true;
      }
    }
  }
  return false;
}

function checkDiago() {
  console.log("D");
  if (!players.includes(document.getElementById(5).value)) {
    return false;
  }
  if (document.getElementById(1).value == document.getElementById(5).value) {
    if (document.getElementById(1).value == document.getElementById(9).value) {
      document.getElementById(1).classList.replace("grid-btn", "grid-btn-red");
      document.getElementById(5).classList.replace("grid-btn", "grid-btn-red");
      document.getElementById(9).classList.replace("grid-btn", "grid-btn-red");
      disable_all();
      if (document.getElementById(1).value == "X") {
        x_score++;
        return true;
      }
      o_score++;
      return true;
    }
  }
  if (document.getElementById(3).value == document.getElementById(5).value) {
    if (document.getElementById(3).value == document.getElementById(7).value) {
      document.getElementById(3).classList.replace("grid-btn", "grid-btn-red");
      document.getElementById(5).classList.replace("grid-btn", "grid-btn-red");
      document.getElementById(7).classList.replace("grid-btn", "grid-btn-red");
      disable_all();
      if (document.getElementById(1).value == "X") {
        x_score++;
        return true;
      }
      o_score++;
      return true;
    }
  }
  return false;
}

function disable_all() {
  for (let i = 1; i <= 9; i++) {
    document.getElementById(i).disabled = true;
  }
}
function reset() {
  for (let i = 1; i <= 9; i++) {
    document.getElementById(i).innerHTML = "_";
    document.getElementById(i).value = "_";
    document.getElementById(i).disabled = false;
    document.getElementById(i).classList.replace("grid-btn-red", "grid-btn");
  }
}
