from tkinter import *
import tkinter as tk
import matplotlib.pyplot as plt
from PIL import  ImageTk, Image
import ExpenseTracker_LogIn as LI
from datetime import datetime
#Window setup
LI.window.destroy()
window = tk.Tk()
window.geometry("600x610")
window.title("ExpenseTracker-MainMenu")
window.config(bg="#343541")
icon = PhotoImage(file="pngwing.com.png")
window.iconphoto(False,icon )
#!!! Fix everything : error on write data if if statement false ; speed; etc!!!
#Variables
bg = "#343541"
data ={}
Error_msg = tk.Variable()
Error_msg.set("")
CurrentView = "Table"
photo_image = None
times_Gen_graph_ran = 0
#funcs
#Checked 6/7/23
def Change_view_mode():
    global CurrentView,Graph_label
    if CurrentView == "Table":
        MM_table_fr.pack_forget()
        Gen_graph()
        CurrentView = "Graph"
    elif CurrentView == "Graph":
        MM_graph_fr.pack_forget()
        Gen_table()
        CurrentView = "Table"
#Checked 6/7/23
def Gen_graph():
    global photo_image,times_Gen_graph_ran
    MM_graph_fr.pack()
    x = []
    y = []
    for key in range(0, len(data)):
        value_list = list(list(data.values())[key])
        date = list(data)[key]
        x.append(date)
    x.sort(key=lambda date: datetime.strptime(date, "%d/%M/%y"))
    for key in x :
        amount = list(data.get(key))[1]
        y.append(int(amount))
    fig = plt.figure()
    ax = fig.add_subplot(111)
    plt.style.use("seaborn")
    plt.plot(x, y, color="#f6c85f", lw=2)
    plt.autoscale(True)
    plt.grid(axis="both")
    ax.tick_params(colors='white')
    plt.setp(ax.spines.values(),color="white")
    plt.savefig('Graph.png', bbox_inches='tight',transparent = True)
    if times_Gen_graph_ran ==0:
        graph_image = Image.open("Graph.png")
        MM_graph_fr.update_idletasks()
        Frame_size = (MM_fr.winfo_width(), MM_fr.winfo_height())
        Resized_graph = graph_image.resize(Frame_size)
        photo_image = ImageTk.PhotoImage(Resized_graph)
        Graph_label.image = Resized_graph
        Graph_label.config(image=photo_image)
    times_Gen_graph_ran +=1
#Checked 6/7/23
def Gen_table():
    global data
    MM_table_fr.pack()
    colors = ["#5A5A5A", "#333333", "#777777","#CACACA"]
    font= "Verdana 17"
    Label(master=MM_table_fr, text="Date",bg=colors[1],font=font,fg="white").grid(row=0,sticky="ew")
    Label(master=MM_table_fr, text="Name",bg=colors[0],font=font,fg="white").grid(row=0,column=1,sticky="ew")
    Label(master=MM_table_fr, text="Amount",bg=colors[1],font=font,fg="white").grid(row=0,column=2,sticky="ew")
    Label(master=MM_table_fr, text="Domain",bg=colors[0],font=font,fg="white").grid(row=0,column=3,sticky="ew")
    total_amount = 0
    for key in range(0,len(data)):
        Label(master=MM_table_fr,text=list(data)[key],bg=colors[2],font=font,fg="white").grid(row=1+key,column=0,sticky="ew")
        value_list = list(list(data.values())[key])
        value_0 = value_list[0]
        value_1 = value_list[1]
        value_2 = value_list[2]
        Label(master=MM_table_fr,text=value_0,bg=colors[3],font=font,fg="black").grid(row=1+key,column=1,sticky="ew")
        Label(master=MM_table_fr,text=value_1,bg=colors[2],font=font,fg="white").grid(row=1+key,column=2,sticky="ew")
        Label(master=MM_table_fr,text=value_2,bg=colors[3],font=font,fg="black").grid(row=1+key,column=3,sticky="ew")
        total_amount += int(value_1)
        Label(master=MM_table_fr, text="Total:", bg=colors[0], font=font, fg="white").grid(row=2 + key, column=1, sticky="ew")
        Label(master=MM_table_fr, text=total_amount, bg=colors[1], font=font, fg="white").grid(row=2 + key, column=2, sticky="ew")

def Save_expense_data():
    global Error_msg
    name = MM_name_fld.get()
    amount = MM_amount_fld.get()
    date = MM_date_fld.get()
    domain = MM_domain_fld.get()
    try:
        amount = int(amount)
    except ValueError:
        Error_msg.set("Amount must be a number")
    data_str ="{};{};{};{}".format(name,amount,date,domain)
    if name != "Name" and amount!="Amount" and date != "Date" and domain!="Domain":
        if (type(amount) == int or type(amount) == float):
            res = None
            try:
                res = bool(datetime.strptime(date, "%d/%M/%y"))
            except:
                Error_msg.set("Required Date format: d/m/yy")
            if res == True:
                with open("ExpenseData.txt","a") as file:
                    Error_msg.set("")
                    file.writelines(data_str +"\n")
                    file.close()
                    Read_expense_data()

        else:
            Error_msg.set("Amount must be a number")

def Read_expense_data():
    global data
    print(data)
    with open("ExpenseData.txt","r") as file:
        Lines = file.readlines()
        for line in Lines:
            data_of_line = line.strip().split(";")
            data[data_of_line[2]] = [data_of_line[0],data_of_line[1],data_of_line[3]]
        print(data)
        file.close()
        Gen_table()

def Bind_FocusIn(field,text):
    field.delete(0,"end")
    field.bind("<FocusOut>",lambda event: field.insert(0, text) if field.get() != text and field.get() == "" else None)

#Widgets for Main menu

MM_fr = Frame(master=window,bg=bg)
MM_Title_fr = Frame(master=MM_fr,bg=bg)
MM_Top_fr = Frame(master=MM_fr,bg=bg)
MM_Top_button_fr =Frame(master=MM_fr,bg=bg)
MM_table_fr = Frame(master=MM_fr,bg=bg,relief=RIDGE,borderwidth=4)
MM_graph_fr = Frame(master=MM_fr,bg=bg)

Graph_label = Label(master=MM_graph_fr, image=photo_image, bg=bg)
Graph_label.pack()


MM_title_lbl = Label(master=MM_Title_fr,bg=bg,font="Helvetica 25",fg= "white",text="Add Expense")
MM_change_view_btn = Button(master=MM_Title_fr,bg="#40414f",font="Helvetica 9",fg= "white",text="Graph",command=Change_view_mode)
MM_Error_lbl = Label(master=MM_fr,bg=bg,font="Calibri 19",fg= "#FF2400",textvariable=Error_msg)

MM_name_fld = Entry(master=MM_Top_fr,bg="#40414f",font="Calibri 19",fg= "white",borderwidth=2)
MM_name_fld.insert(0,"Name")
MM_name_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(MM_name_fld,"Name"))

MM_amount_fld = Entry(master=MM_Top_fr,bg="#40414f",font="Calibri 19",fg= "white",borderwidth=2)
MM_amount_fld.insert(0,"Amount")
MM_amount_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(MM_amount_fld,"Amount"))

MM_date_fld = Entry(master=MM_Top_fr,bg="#40414f",font="Calibri 19",fg= "white",borderwidth=2)
MM_date_fld.insert(0,"Date")
MM_date_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(MM_date_fld,"Date"))

MM_domain_fld = Entry(master=MM_Top_fr,bg="#40414f",font="Calibri 19",fg= "white",borderwidth=2)
MM_domain_fld.insert(0,"Domain")
MM_domain_fld.bind("<FocusIn>",lambda event: Bind_FocusIn(MM_domain_fld,"Domain"))
MM_Add_exp_btn = Button(master=MM_Top_button_fr,bg="#40414f",font="Helvetica 11",fg= "white",text="Add",command=Save_expense_data)
#Packing for Main Menu
MM_title_lbl.pack(side="left",padx=(0,325),pady=15,anchor="w")
MM_Error_lbl.place(x=100,y=65)
MM_change_view_btn.pack(side="left",anchor="e")

MM_name_fld.grid(row=1,column=1,pady=(40,0),ipady=5)
MM_amount_fld.grid(row=1,column=2,padx=(20,0),pady=(40,0),ipady=5)
MM_date_fld.grid(row=2,column=1,pady=(40,0),ipady=5)
MM_domain_fld.grid(row=2,column=2,padx=(20,0),pady=(40,0),ipady=5)
MM_Add_exp_btn.pack(side="right",pady=10,padx=(514,0))


MM_Title_fr.pack(side="top")
MM_Top_fr.pack(side="top")
MM_Top_button_fr.pack(side="top")
MM_table_fr.pack(side="top")


MM_fr.pack()
Read_expense_data()
#MainLoop
window.mainloop()
