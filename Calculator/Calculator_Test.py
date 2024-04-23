from tkinter import *
import tkinter as tk
from math import sqrt
import keyboard
window = tk.Tk()
icon_photo = PhotoImage(file="Windows_Calculator_icon.png")
window.iconphoto(False,icon_photo)
window.geometry("314x455")
window.title("Calculator")
window.resizable(False,False)
window.config(bg="#202020")
bg_standard = "#202020"
lower_button_bg ="#3c3c3c"
high_button_bg = "#323232"
equal_button_bg = "#dd765d"
Result_tkvar = tk.Variable(value=0)
Expression_tkvar = tk.Variable(value="")
expression_final = ""
result = ""
expression_temporary =""
last_function   = ""
def C_delete():
    global expression_temporary,expression_final,result,last_function
    Result_tkvar.set("0")
    Output_label_bottom.config(font="Verdana 35")
    Input_frame.pack(side="top")
    Expression_tkvar.set("")
    expression_final = ""
    result = ""
    expression_temporary = ""
    last_function = ""
def Percent():#https://devblogs.microsoft.com/oldnewthing/20080110-00/?p=23853
    global expression_temporary, last_function,expression_final
    if last_function == "add_num":
        expression_final = expression_final
        if "+" in expression_final: expression_final = expression_final.replace("+","")
        if "-" in expression_final: expression_final = expression_final.replace("-","")
        if "/" in expression_final: expression_final = expression_final.replace("/","")
        if "*" in expression_final: expression_final_Pe55rcent = expression_final.replace("*","")
        expression_temporary = str(int(expression_temporary)/100)
        try:
            expression_final = float(expression_final) + float(expression_final) * float(expression_temporary)
            Result_tkvar.set(expression_final)
        except ValueError:
            Result_tkvar.set("0")

def Inverse():
    global expression_temporary,last_function
    if last_function == "add_num":
        Expression_tkvar.set("1/(" + expression_temporary + ")")
        expression_temporary = str(1/float(expression_temporary))
        if len(str(expression_temporary)) >10 :
            expression_temporary = str(expression_temporary)[:10]
        Result_tkvar.set(expression_temporary)
def Square():
    global expression_temporary,last_function
    if last_function == "add_num":
        Expression_tkvar.set("sqr(" + expression_temporary + ")")
        expression_temporary = str(int(expression_temporary)*int(expression_temporary))
        Result_tkvar.set(expression_temporary)
def Delete():
    global expression_temporary,last_function
    if last_function == "add_num" :
        expression_temporary = list(expression_temporary)
        expression_temporary.pop()
        other_exp_temp = ""
        for i in expression_temporary:
            other_exp_temp += str(i)
            print(other_exp_temp)
        expression_temporary = str(other_exp_temp)
        Result_tkvar.set(expression_temporary)
        last_function = "add_num"
def Dot():
    global expression_temporary, last_function, times_dot_used
    if last_function == "add_num" and "." not in expression_temporary:
        # Expression_tkvar.set("1/(" + expression_temporary + ")")
        expression_temporary += "."
        print(expression_temporary)
        Result_tkvar.set(expression_temporary)
def CE_delete():
    global expression_temporary, last_function
    if last_function == "add_num":
        # Expression_tkvar.set("1/(" + expression_temporary + ")")
        expression_temporary = ""
        Result_tkvar.set("0")
def Square_root():
    global expression_temporary, last_function
    if last_function == "add_num":
        Expression_tkvar.set("sqrt(" + expression_temporary + ")")
        expression_temporary = str(sqrt(float(expression_temporary)))
        if len(str(expression_temporary)) >10 :
            expression_temporary = str(expression_temporary)[:10]
        Result_tkvar.set(expression_temporary)
def Opposite():
    global expression_temporary,last_function
    if last_function == "add_num":
        expression_temporary = str(int(expression_temporary)*-1)
        print(expression_temporary)
        Result_tkvar.set(expression_temporary)
def Add_number_to_expression(num):
    global expression_temporary,expression_final,last_function
    if last_function == "add_op":
        expression_final += expression_temporary
        expression_temporary = ""
        print("num-op",expression_final)
    expression_temporary += num
    Result_tkvar.set(f"{float(expression_temporary) :,}")
    if last_function == "":Expression_tkvar.set(expression_temporary)
    else:Expression_tkvar.set(expression_final)
    last_function = "add_num"
def Add_operation_to_expression(op):
    global expression_temporary,expression_final,last_function,result
    if last_function == "add_op":
        expression_temporary = op
        Expression_tkvar.set(expression_final + expression_temporary)
        print("op-num",expression_final)
    elif last_function == "equals":
        expression_final  = str(result)
        result = ''
        expression_temporary = op
        Expression_tkvar.set(expression_final+op)
        last_function ="add_op"
    elif (last_function == "add_num") and (expression_final.count("+") >=1 or expression_final.count("-") >=1 or expression_final.count("*") >=1 or expression_final.count("/") >=1) :
        result = str(eval(expression_final +expression_temporary))
        expression_final = result
        print(result)
        expression_temporary = op
        Expression_tkvar.set(expression_final + op)
        last_function = "add_op"
    else:
        expression_final += expression_temporary
        expression_temporary = op
        print(expression_final.count("+"))
        Expression_tkvar.set(expression_final +op)
        last_function = "add_op"
def Equals():
    global result,expression_final,last_function
    if last_function == "add_num":
        expression_final += expression_temporary
    try:
        result = eval(expression_final)
        if len(str(result)) >10 :
            result = str(result)[:10]
        result = f"{float(result):,}"
        Expression_tkvar.set(expression_final + "=")
        Result_tkvar.set(result)
        last_function = "equals"
    except ZeroDivisionError:
        Result_tkvar.set("Cannot divide by 0")
        Output_label_bottom.config(font="Verdana 20")
        Input_frame.pack(side="top", pady=(35,0))
    except SyntaxError:
        Result_tkvar.set("Try again")
        Output_label_bottom.config(font="Verdana 20")
        Input_frame.pack(side="top", pady=(35,0))

def Detect_key_press(event):

    operators = ['+', '-', '*', '/']
    if event.keysym == "equal":
        Equals()
    if event.keysym == "Return":
        Equals()
    if event.keysym == "Delete":
        Delete()
    if keyboard.is_pressed("shift+5"):
        Percent()
    if keyboard.is_pressed("shift+2"):
        Square_root()
    if keyboard.is_pressed("shift+8"):
        Add_operation_to_expression("*")
    if event.char.isdigit():
        Add_number_to_expression(str(event.char))
    if event.char in operators:
        Add_operation_to_expression(str(event.char))
    if event.keysym == 'Escape':
        C_delete()
    if event.keysym == 'BackSpace':
        Delete()
    if event.char == '.':
        Dot()
Output_frame = Frame(master=window,bg=bg_standard)
Output_label_top = Label(master=Output_frame, bg=bg_standard, textvariable=Expression_tkvar, font="Verdana 14", fg="#7d7d7d", anchor='e', justify=RIGHT)
Output_label_bottom = Label(master=Output_frame, bg=bg_standard, textvariable=Result_tkvar, font="Verdana 35", fg="#e8eaed", anchor='e', justify=RIGHT)

Input_frame = Frame(master=window,bg=bg_standard)
#4X6
Percent_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="%",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Percent)
Percent_button.grid(row=0,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Percent_button.bind("<Enter>",lambda event:Percent_button.config(bg=lower_button_bg))
Percent_button.bind("<Leave>",lambda event:Percent_button.config(bg=high_button_bg))

CE_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="CE",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=CE_delete)
CE_button.grid(row=0,column=1,sticky="ew",padx=2,pady=2,ipady=3)
CE_button.bind("<Enter>",lambda event:CE_button.config(bg=lower_button_bg))
CE_button.bind("<Leave>",lambda event:CE_button.config(bg=high_button_bg))

C_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="C",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=C_delete)
C_button.grid(row=0,column=2,sticky="ew",padx=2,pady=2,ipady=3)
C_button.bind("<Enter>",lambda event:C_button.config(bg=lower_button_bg))
C_button.bind("<Leave>",lambda event:C_button.config(bg=high_button_bg))

Delete_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="⌫",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Delete)
Delete_button.grid(row=0,column=3,sticky="ew",padx=2,pady=2,ipady=3)
Delete_button.bind("<Enter>",lambda event:Delete_button.config(bg=lower_button_bg))
Delete_button.bind("<Leave>",lambda event:Delete_button.config(bg=high_button_bg))

Inverse_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="¹/ₓ ",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Inverse)
Inverse_button.grid(row=1,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Inverse_button.bind("<Enter>",lambda event:Inverse_button.config(bg=lower_button_bg))
Inverse_button.bind("<Leave>",lambda event:Inverse_button.config(bg=high_button_bg))
Square_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="x²",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Square)
Square_button.grid(row=1,column=1,sticky="ew",padx=2,pady=2,ipady=3)
Square_button.bind("<Enter>",lambda event:Square_button.config(bg=lower_button_bg))
Square_button.bind("<Leave>",lambda event:Square_button.config(bg=high_button_bg))

SquareRoot_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="√x",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Square_root)
SquareRoot_button.grid(row=1,column=2,sticky="ew",padx=2,pady=2,ipady=3)
SquareRoot_button.bind("<Enter>",lambda event:SquareRoot_button.config(bg=lower_button_bg))
SquareRoot_button.bind("<Leave>",lambda event:SquareRoot_button.config(bg=high_button_bg))

Divide_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="÷",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda: Add_operation_to_expression("/"))
Divide_button.grid(row=1,column=3,sticky="ew",padx=2,pady=2,ipady=3)
Divide_button.bind("<Enter>",lambda event:Divide_button.config(bg=lower_button_bg))
Divide_button.bind("<Leave>",lambda event:Divide_button.config(bg=high_button_bg))

Num7_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="7",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("7"))
Num7_button.grid(row=2,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Num7_button.bind("<Enter>",lambda event:Num7_button.config(bg=high_button_bg))
Num7_button.bind("<Leave>",lambda event:Num7_button.config(bg=lower_button_bg))

Num8_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="8",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("8"))
Num8_button.grid(row=2,column=1,sticky="ew",padx=2,pady=2,ipady=3)
Num8_button.bind("<Enter>",lambda event:Num8_button.config(bg=high_button_bg))
Num8_button.bind("<Leave>",lambda event:Num8_button.config(bg=lower_button_bg))

Num9_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="9",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("9"))
Num9_button.grid(row=2,column=2,sticky="ew",padx=2,pady=2,ipady=3)
Num9_button.bind("<Enter>",lambda event:Num9_button.config(bg=high_button_bg))
Num9_button.bind("<Leave>",lambda event:Num9_button.config(bg=lower_button_bg))

Multiplication_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="x",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_operation_to_expression("*"))
Multiplication_button.grid(row=2,column=3,sticky="ew",padx=2,pady=2,ipady=3)
Multiplication_button.bind("<Enter>",lambda event:Multiplication_button.config(bg=lower_button_bg))
Multiplication_button.bind("<Leave>",lambda event:Multiplication_button.config(bg=high_button_bg))

Num4_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="4",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("4"))
Num4_button.grid(row=3,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Num4_button.bind("<Enter>",lambda event:Num4_button.config(bg=high_button_bg))
Num4_button.bind("<Leave>",lambda event:Num4_button.config(bg=lower_button_bg))

Num5_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="5",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("5"))
Num5_button.grid(row=3,column=1,sticky="ew",padx=2,pady=2,ipady=3)
Num5_button.bind("<Enter>",lambda event:Num5_button.config(bg=high_button_bg))
Num5_button.bind("<Leave>",lambda event:Num5_button.config(bg=lower_button_bg))

Num6_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="6",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("6"))
Num6_button.grid(row=3,column=2,sticky="ew",padx=2,pady=2,ipady=3)
Num6_button.bind("<Enter>",lambda event:Num6_button.config(bg=high_button_bg))
Num6_button.bind("<Leave>",lambda event:Num6_button.config(bg=lower_button_bg))

Subtraction_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="-",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda: Add_operation_to_expression("-"))
Subtraction_button.grid(row=3,column=3,sticky="ew",padx=2,pady=2,ipady=3)
Subtraction_button.bind("<Enter>",lambda event:Subtraction_button.config(bg=lower_button_bg))
Subtraction_button.bind("<Leave>",lambda event:Subtraction_button.config(bg=high_button_bg))

Num1_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="1",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("1"))
Num1_button.grid(row=4,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Num1_button.bind("<Enter>",lambda event:Num1_button.config(bg=high_button_bg))
Num1_button.bind("<Leave>",lambda event:Num1_button.config(bg=lower_button_bg))

Num2_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="2",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("2"))
Num2_button.grid(row=4,column=1,sticky="ew",padx=2,pady=2,ipady=3)
Num2_button.bind("<Enter>",lambda event:Num2_button.config(bg=high_button_bg))
Num2_button.bind("<Leave>",lambda event:Num2_button.config(bg=lower_button_bg))

Num3_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="3",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("3"))
Num3_button.grid(row=4,column=2,sticky="ew",padx=2,pady=2,ipady=3)
Num3_button.bind("<Enter>",lambda event:Num3_button.config(bg=high_button_bg))
Num3_button.bind("<Leave>",lambda event:Num3_button.config(bg=lower_button_bg))

Additon_button = Button(master=Input_frame,bg=high_button_bg,activebackground=lower_button_bg,text="+",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda: Add_operation_to_expression("+"))
Additon_button.grid(row=4,column=3,sticky="ew",padx=2,pady=2,ipady=3)
Additon_button.bind("<Enter>",lambda event:Additon_button.config(bg=lower_button_bg))
Additon_button.bind("<Leave>",lambda event:Additon_button.config(bg=high_button_bg))

Opposite_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="±",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Opposite)
Opposite_button.grid(row=5,column=0,sticky="ew",padx=2,pady=2,ipady=3)
Opposite_button.bind("<Enter>",lambda event:Opposite_button.config(bg=high_button_bg))
Opposite_button.bind("<Leave>",lambda event:Opposite_button.config(bg=lower_button_bg))

Num0_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text="0",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=lambda:Add_number_to_expression("0"))
Num0_button.grid(row=5,column=1,sticky="ew",padx=2,pady=2,ipady=3)
Num0_button.bind("<Enter>",lambda event:Num0_button.config(bg=high_button_bg))
Num0_button.bind("<Leave>",lambda event:Num0_button.config(bg=lower_button_bg))

Dot_button = Button(master=Input_frame,bg=lower_button_bg,activebackground=high_button_bg,text=".",borderwidth=0,relief=FLAT,activeforeground="white",font="Helvetica  17",fg="white",width=5,command=Dot)
Dot_button.grid(row=5,column=2,sticky="ew",padx=2,pady=2,ipady=3)
Dot_button.bind("<Enter>",lambda event:Dot_button.config(bg=high_button_bg))
Dot_button.bind("<Leave>",lambda event:Dot_button.config(bg=lower_button_bg))

Equal_button = Button(master=Input_frame, bg=equal_button_bg, text="=", borderwidth=0, relief=FLAT, activebackground="#c86c56", activeforeground="#894a3b", font="Helvetica  17", fg="white", width=5, command=Equals)
Equal_button.grid(row=5,column=3,sticky="ew",padx=2,pady=2,ipady=3)
window.bind("<KeyPress>",Detect_key_press)

Output_frame.pack(side="top")
Output_label_top.pack(side="top",fill="x",ipadx=200,pady=(20,0),padx=(0,10))
Output_label_bottom.pack(side="top",fill="x",ipadx=200,pady=(0,20))
Input_frame.pack(side="top")
window.mainloop()