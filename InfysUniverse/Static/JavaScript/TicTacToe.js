let Curr_Player = 0;
let x_score = 0;
let o_score = 0;
function setButton(id) {
  let header = document.getElementById(id);
  if (header.value == "_") {
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
  CheckEnd();
}
function CheckEnd() {
  const horizontal = [1, 4, 7];
  const vertical = [1, 2, 3];
  for (let index in horizontal) {
    const i = horizontal[index];
    console.log(i);
    if (document.getElementById(i).value != "_") {
      if (
        document.getElementById(i).value == document.getElementById(i + 1).value
      ) {
        if (
          document.getElementById(i).value ==
          document.getElementById(i + 2).value
        ) {
          disable_all();
          if (document.getElementById(i).value == "X") {
            x_score++;

            document.getElementById("x_sc").innerHTML = x_score;
          } else {
            o_score++;
            document.getElementById("o_sc").innerHTML = o_score;
          }
        }
      }
    }
  }
  for (let i in vertical) {
    const num = vertical[i];
    if (document.getElementById(num).value != "_") {
      if (
        document.getElementById(num).value ==
        document.getElementById(num + 3).value
      ) {
        if (
          document.getElementById(num).value ==
          document.getElementById(num + 6).value
        ) {
          disable_all();
          if (document.getElementById(num).value == "X") {
            x_score++;
            document.getElementById("x_sc").innerHTML = x_score;
          } else {
            o_score++;
            document.getElementById("o_sc").innerHTML = o_score;
          }
        }
      }
    }
  }

  if (document.getElementById(1).value != "_") {
    if (document.getElementById(1).value == document.getElementById(5).value) {
      if (
        document.getElementById(1).value == document.getElementById(9).value
      ) {
        disable_all();
        if (document.getElementById(1).value == "X") {
          x_score++;
          document.getElementById("x_sc").innerHTML = x_score;
        } else {
          o_score++;
          document.getElementById("o_sc").innerHTML = o_score;
        }
      }
    }
  }
  if (document.getElementById(3).value != "_") {
    if (document.getElementById(3).value == document.getElementById(5).value) {
      if (
        document.getElementById(3).value == document.getElementById(7).value
      ) {
        disable_all();
        if (document.getElementById(3).value == "X") {
          x_score++;
          document.getElementById("x_sc").innerHTML = x_score;
        } else {
          o_score++;
          document.getElementById("o_sc").innerHTML = o_score;
        }
      }
    }
  }
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
  }
}
