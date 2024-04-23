from tkinter import *
import tkinter as tk

window = tk.Tk()
window.geometry("550x620")
window.config(bg="#14bdac")
window.title("TicTacToe")
window.resizable(False, False)
bg_color = "#14bdac"
button_bg = "#0da192"

X_Score = tk.Variable()
O_Score = tk.Variable()
X_Score.set(0)
O_Score.set(0)
Current_player = "x"
Button_variable_1 = tk.Variable()
Button_variable_2 = tk.Variable()
Button_variable_3 = tk.Variable()
Button_variable_4 = tk.Variable()
Button_variable_5 = tk.Variable()
Button_variable_6 = tk.Variable()
Button_variable_7 = tk.Variable()
Button_variable_8 = tk.Variable()
Button_variable_9 = tk.Variable()
Button_color_1 = tk.Variable()
Button_color_2 = tk.Variable()
Button_color_3 = tk.Variable()
Button_color_4 = tk.Variable()
Button_color_5 = tk.Variable()
Button_color_6 = tk.Variable()
Button_color_7 = tk.Variable()
Button_color_8 = tk.Variable()
Button_color_9 = tk.Variable()
Winner = ""
Win_Label = Label(master=window, text="", bg=bg_color, font="Helvetica 25")


def Check_win():
    global Winner, X_Score, O_Score, Win_Label
    # H
    if (
        (Button_variable_1.get() == Button_variable_2.get())
        and (Button_variable_3.get() == Button_variable_1.get())
        and (Button_variable_1.get() != "")
    ):
        Winner = Button_variable_3.get()
        Button_color_3.set("red")
        Button_color_1.set("red")
        Button_color_2.set("red")
    elif (
        (Button_variable_4.get() == Button_variable_5.get())
        and (Button_variable_4.get() == Button_variable_6.get())
        and (Button_variable_4.get() != "")
    ):
        Winner = Button_variable_6.get()
        Button_color_4.set("red")
        Button_color_5.set("red")
        Button_color_6.set("red")
    elif (
        (Button_variable_7.get() == Button_variable_8.get())
        and (Button_variable_7.get() == Button_variable_9.get())
        and (Button_variable_7.get() != "")
    ):
        Winner = Button_variable_9.get()
        Button_color_7.set("red")
        Button_color_8.set("red")
        Button_color_9.set("red")
    # V
    elif (
        (Button_variable_1.get() == Button_variable_4.get())
        and (Button_variable_7.get() == Button_variable_1.get())
        and (Button_variable_1.get() != "")
    ):
        Winner = Button_variable_1.get()
        Button_color_1.set("red")
        Button_color_4.set("red")
        Button_color_7.set("red")
    elif (
        (Button_variable_2.get() == Button_variable_5.get())
        and (Button_variable_8.get() == Button_variable_2.get())
        and (Button_variable_2.get() != "")
    ):
        Winner = Button_variable_2.get()
        Button_color_8.set("red")
        Button_color_5.set("red")
        Button_color_2.set("red")
    elif (
        (Button_variable_3.get() == Button_variable_6.get())
        and (Button_variable_9.get() == Button_variable_3.get())
        and (Button_variable_3.get() != "")
    ):
        Winner = Button_variable_3.get()
        Button_color_3.set("red")
        Button_color_6.set("red")
        Button_color_9.set("red")
    # D
    elif (
        (Button_variable_1.get() == Button_variable_5.get())
        and (Button_variable_9.get() == Button_variable_1.get())
        and (Button_variable_1.get() != "")
    ):
        Winner = Button_variable_1.get()
        Button_color_5.set("red")
        Button_color_1.set("red")
        Button_color_9.set("red")
    elif (
        (Button_variable_3.get() == Button_variable_5.get())
        and (Button_variable_7.get() == Button_variable_3.get())
        and (Button_variable_3.get() != "")
    ):
        Winner = Button_variable_3.get()
        Button_color_3.set("red")
        Button_color_5.set("red")
        Button_color_7.set("red")

    if (
        Button_variable_1.get() != ""
        and Button_variable_2.get() != ""
        and Button_variable_3.get() != ""
        and Button_variable_4.get() != ""
        and Button_variable_5.get() != ""
        and Button_variable_6.get() != ""
        and Button_variable_7.get() != ""
        and Button_variable_8.get() != ""
        and Button_variable_9.get() != ""
        and Winner == ""
    ):
        Winner = "Draw"
    if Winner != "" and Winner != "Draw":
        Win_Label = Label(
            master=window, text="Winner: " + Winner, bg=bg_color, font="Helvetica 25"
        )
        Win_Label.place(rely=0.05, relx=0.4)
        Create_grid_v2()
        if Winner == "X":
            X_Score.set(X_Score.get() + 1)
        else:
            O_Score.set(O_Score.get() + 1)
    elif Winner == "Draw":
        Win_Label = Label(master=window, text="Draw", bg=bg_color, font="Helvetica 25")
        Win_Label.place(rely=0.05, relx=0.4)
        Create_grid_v2()


def Button_press(Button_variable):
    global Current_player
    if Button_variable.get() != "X" and Button_variable.get() != "O" and Winner == "":
        if Current_player == "x":
            Button_variable.set("X")
            Current_player = "o"
            Check_win()
        elif Current_player == "o":
            Button_variable.set("O")
            Current_player = "x"
            Check_win()


def Create_grid_v2():
    Button_row_column_dictionary = {
        1: ("row=0,column=0", Button_variable_1, Button_color_1),
        2: ("row=0,column=1", Button_variable_2, Button_color_2),
        3: ("row=0,column=2", Button_variable_3, Button_color_3),
        4: ("row=1,column=0", Button_variable_4, Button_color_4),
        5: ("row=1,column=1", Button_variable_5, Button_color_5),
        6: ("row=1,column=2", Button_variable_6, Button_color_6),
        7: ("row=2,column=0", Button_variable_7, Button_color_7),
        8: ("row=2,column=1", Button_variable_8, Button_color_8),
        9: ("row=2,column=2", Button_variable_9, Button_color_9),
    }
    for key in range(0, 9):
        row_column = list(Button_row_column_dictionary.values())[key][0]
        command = list(Button_row_column_dictionary.values())[key][1]
        collor = list(Button_row_column_dictionary.values())[key][2]
        eval(
            """Button(master=Game_frame,
                relief=RIDGE,
                width=3,
                height=1,
                activebackground=button_bg,
                bg=collor.get(),
                borderwidth=5,
                textvariable= command,
                font = "Verdana 50",
                command=lambda cmd=command:Button_press(cmd),
                overrelief = GROOVE).grid("""
            + row_column
            + ")"
        )


def Reset_game():
    global X_Score, O_Score, Win_Label, Current_player, Winner
    Win_Label.place_forget()
    Current_player = "x"
    Button_variable_1.set("")
    Button_variable_2.set("")
    Button_variable_3.set("")
    Button_variable_4.set("")
    Button_variable_5.set("")
    Button_variable_6.set("")
    Button_variable_7.set("")
    Button_variable_8.set("")
    Button_variable_9.set("")
    Button_color_1.set(bg_color)
    Button_color_2.set(bg_color)
    Button_color_3.set(bg_color)
    Button_color_4.set(bg_color)
    Button_color_5.set(bg_color)
    Button_color_6.set(bg_color)
    Button_color_7.set(bg_color)
    Button_color_8.set(bg_color)
    Button_color_9.set(bg_color)
    Winner = ""
    Create_grid_v2()


General_canvas = Canvas(
    master=window,
    bg=bg_color,
    borderwidth=0,
    highlightthickness=0,
    width=550,
    height=620,
)
Score_frame = Frame(master=General_canvas, bg=bg_color)
Game_frame = Frame(master=General_canvas, bg=bg_color)
X_Score_Title_label = Label(
    master=Score_frame, bg=bg_color, text="X: ", font="Helvetica 25 bold"
)
O_Score_Title_label = Label(
    master=Score_frame, bg=bg_color, text="O: ", font="Helvetica 25 bold"
)

X_Score_label = Label(
    master=Score_frame, bg=bg_color, textvariable=X_Score, font="Calibri 25 "
)
O_Score_label = Label(
    master=Score_frame, bg=bg_color, textvariable=O_Score, font="Calibri 25 "
)

Reset_button = Button(
    master=Score_frame,
    bg=bg_color,
    text="Re",
    command=Reset_game,
    font="Calibri 10 bold",
    borderwidth=4,
)
X_Score_Title_label.pack(pady=(80, 0), side="left", padx=(550, 0))
X_Score_label.pack(pady=(80, 0), side="left")
O_Score_Title_label.pack(pady=(80, 0), side="left", padx=(300, 0))
O_Score_label.pack(pady=(80, 0), side="left")
Reset_button.pack(side="left", pady=(80, 0), padx=(10, 0))
Score_win = General_canvas.create_window(0, 0, window=Score_frame)
Game_win = General_canvas.create_window(270, 340, window=Game_frame)


General_canvas.pack()
Reset_game()

window.mainloop()
