import sys
from tkinter import *
import tkinter as tk
import os
from secrets import compare_digest
#Window setup
window = tk.Tk()
window.geometry("150x220")
window.title("ExpenseTracker-LogIn")
window.config(bg="#343541")
window.resizable(False,False)
icon = PhotoImage(file="pngwing.com.png")
window.iconphoto(False,icon)


#Variables
Inv_cred = tk.Variable()

#funcs
def Bind_FocusIn(field,text):
    field.delete(0,"end")
    field.bind("<FocusOut>", lambda event: field.insert(0, text) if field.get() != text and field.get() == "" else None)


def Log_In():
    Inp_user = LI_User_fld.get()
    Inp_pass = LI_Pass_fld.get()
    Txt_to_write = "{} {}".format(Inp_user, Inp_pass)
    with open("Credentials.txt", "r") as file:
        line = file.readline()
        line_arr = line.strip().split()
        if len(line_arr) == 2:
            Inv_cred.set("")
            r_user = line_arr[0]
            r_pass = line_arr[1]
            if compare_digest(Inp_user, r_user) and compare_digest(Inp_pass, r_pass):
                Inv_cred.set("")
                Error_lbl.pack_forget()
                import ExpenseTracker_MainMenu as MM
                window.pack_forget()
            else:
                Inv_cred.set("Invalid Credentials")
        else:
            Inv_cred.set("Error credentials <2")

def Create():
    Inp_user = LI_User_fld.get()
    Inp_pass = LI_Pass_fld.get()
    Txt_to_write = "{} {}".format(Inp_user, Inp_pass)
    with open("Credentials.txt", "a+") as file:
        Line = file.readline()
        if os.stat("Credentials.txt").st_size == 0:
            file.write(Txt_to_write)
        else:
            Inv_cred.set("Account already exists")

#Widgets
LI_fr= Frame(master=window,bg="#343541")
LI_lbl = Label(master=LI_fr,bg="#343541",fg="white",text="Welcome Back",font="OpenSans 12 bold")
Error_lbl = Label(master=LI_fr,bg="#343541",fg= "#FF2400",textvariable=Inv_cred)
#4C4C59
LI_User_fld = Entry(master=LI_fr,bg="#40414f",fg="white",borderwidth=2)
LI_User_fld.insert(0,"User Name")
LI_User_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(LI_User_fld,"User Name"))

LI_Pass_fld = Entry(master=LI_fr,bg="#40414f",fg="white",borderwidth=2)
LI_Pass_fld.insert(0,"Password")

LI_Pass_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(LI_Pass_fld,"Password"))

LI_Log_btn = Button(master=LI_fr,bg="#40414f",borderwidth=2,fg="white",text="Log In",font="OpenSans 11 ",command=Log_In)
LI_create_btn = Button(master=LI_fr,bg="#40414f",borderwidth=2,fg="white",text="Create Account",font="OpenSans 8 ",command=Create)




LI_lbl.pack()
Error_lbl.pack()
LI_User_fld.pack(pady=5)
LI_Pass_fld.pack(pady=10)
LI_Log_btn.pack(pady=10)
LI_create_btn.pack(pady=10)
LI_fr.pack(pady=10)
#MainLoop
window.mainloop()