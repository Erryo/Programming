from tkinter import *
import tkinter as tk
import Keyboard as kb
window = tk.Tk()
window.geometry("600x600")
window.resizable(False,False)
window.title("Hangman")
window.config(bg="#6cc5ce")
Word_given = tk.Variable(value="")
Used_letters = []
Letters_of_word ={}
Copy_word =""
Hidden_word = tk.Variable(value=Copy_word)
def Restart():
    global Used_letters,Letters_of_word,Copy_word
    Word_given.set(value="")
    Used_letters = []
    Letters_of_word = {}
    Copy_word = ""
    Hidden_word.set(value=Copy_word)
    Win_label.place_forget()
    Lose_label.place_forget()
    for i in kb.ButtonName_Letter:
        kb.ButtonName_Letter[i].config(state="disabled", bg="#9ed0d4", relief="flat")
    Secret_word_hidden_label.pack(side="top", anchor="center", fill="both", pady=(30, 0))
    Word_entry.pack()
    Word_entry.config(state="normal")
    Start_frame.place(rely=.5, relx=.05)
    window.bind("<KeyPress>", Secret_word_given)
def Check_letter(argument):
    global Copy_word
    argument = argument.lower()
    if argument not in Used_letters:
        Used_letters.append(argument)
        kb.ButtonName_Letter[argument].config(state="disabled",bg="#9ed0d4",relief="flat")
        if len(Used_letters) <= 6:
            if argument in str(Word_given.get()).lower():
                Letters_of_word[argument] = True
                temp = list(Word_given.get())
                Copy_word=list(Copy_word)
                for char in Word_given.get():
                    if char == argument:
                        for j in [i for i, n in enumerate(temp) if n == char]:
                            Copy_word.pop(j)
                            Copy_word.insert(j, char)
                        Hidden_word.set(Copy_word)
        else:
            Lose_label.place(rely=.15,relx=.35)
            for i in kb.ButtonName_Letter:
                kb.ButtonName_Letter[i].config(state="disabled", bg="#9ed0d4", relief="flat")
    won = True
    test_val = list(Letters_of_word.values())[0]
    for key in Letters_of_word:
        if test_val != Letters_of_word[key] and test_val != False:
            won = False
            break
    if won == True  and test_val != False:
        for i in kb.ButtonName_Letter:
            kb.ButtonName_Letter[i].config(state="disabled", bg="#9ed0d4", relief="flat")
        Win_label.place(rely=.15,relx=.35)

def Secret_word_given(event):
    global Copy_word

    if event.keysym == "Return":
        if Word_given.get() != "":
            if str(Word_given.get()).isalpha():
                Warning_label.place_forget()
                Word_entry.pack_forget()
                Word_entry.config(state="disabled")
                Start_frame.place_forget()
                for char in Word_given.get():
                    Letters_of_word[char] = False
                    Copy_word += "_"
                Hidden_word.set(Copy_word)
                window.bind("<KeyPress>",lambda event:Check_letter(event.keysym) if event.char.isalpha() else print(""))
                for i in kb.ButtonName_Letter:kb.ButtonName_Letter[i].config(state="normal", bg="#f3d611", relief="raised",activebackground="#d8bd0f")
                Reset_button.place(rely=.05, relx=.9)
            else:
                Warning_label.place(relx=.2,rely=.2)
Keyboard_Frame = Frame(master=window,bg="#6cc5ce",relief=RIDGE,borderwidth=4)
Start_frame = Frame(master=window,bg ="#6cc5ce")
Word_entry = Entry(master=Start_frame,textvariable=Word_given,bg="#272838",fg="#F9F8F8",font="Helvetica 36",justify="center")
Warning_label = Label(master=window, text="Insert only words", font="Helvetica 35", fg="Red", bg="#6cc5ce")
kb.Draw_keyboard(master=Keyboard_Frame, width=8, height=4, background="#f3d611", foreground="black",Filler_func=Check_letter,window=window,pady=(2,2),padx=(2,2),borderwidth=2,activebg="#d8bd0f")
Reset_button =Button(master=window,text="↺",bg="#437b81",fg="white",borderwidth=2,font="Verdana 15",relief=RIDGE,overrelief=GROOVE,activebackground="#6cc5ce",command=Restart)
Lose_label=Label(master=window,text="You lost",fg="black",bg="#6cc5ce",font="Helvetica 35")
Win_label=Label(master=window,text="Correct!",fg="black",bg="#6cc5ce",font="Helvetica 35")
Game_frame = Frame(master=window,bg="#6cc5ce")
Game_frame.pack(side="top")
Keyboard_Frame.pack(side="top",pady=(120,0))
Secret_word_hidden_label = Label(master=Game_frame,textvariable=Hidden_word,bg="#6cc5ce",fg="#F9F8F8",font="Helvetica 36 bold")

for i in kb.ButtonName_Letter:
    kb.ButtonName_Letter[i].config(state="disabled", bg="#9ed0d4", relief="flat")
Secret_word_hidden_label.pack(side="top",anchor="center",fill="both",pady=(30,0))
window.bind("<KeyPress>",Secret_word_given)

Word_entry.pack()
Start_frame.place(rely=.5,relx=.05)
window.mainloop()
