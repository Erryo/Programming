from tkinter import *
import tkinter as tk
ButtonName_Letter ={}
def Draw_keyboard(master,Filler_func,window,width=1,height=1,background="white",foreground="black",pady=(0,0),padx=(0,0),borderwidth=0,activebg="white"):
    global ButtonName_Letter
    Letter_a = Button(master=master,bg=background,fg=foreground,text="A",command=lambda:Filler_func("A"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_a.grid(row=0,column=0,sticky="ew",pady=pady,padx=padx)
    Letter_b = Button(master=master,bg=background,fg=foreground,text="B",command=lambda:Filler_func("B"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_b.grid(row=0,column=1,sticky="ew",pady=pady,padx=padx)

    Letter_c = Button(master=master,bg=background,fg=foreground,text="C",command=lambda:Filler_func("C"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_c.grid(row=0,column=2,sticky="ew",pady=pady,padx=padx)

    Letter_d = Button(master=master,bg=background,fg=foreground,text="D",command=lambda:Filler_func("D"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_d.grid(row=0,column=3,sticky="ew",pady=pady,padx=padx)

    Letter_e = Button(master=master,bg=background,fg=foreground,text="E",command=lambda:Filler_func("E"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_e.grid(row=0,column=4,sticky="ew",pady=pady,padx=padx)

    Letter_f = Button(master=master,bg=background,fg=foreground,text="F",command=lambda:Filler_func("F"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_f.grid(row=0,column=5,sticky="ew",pady=pady,padx=padx)


    Letter_g = Button(master=master,bg=background,fg=foreground,text="G",command=lambda:Filler_func("G"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_g.grid(row=1,column=0,sticky="ew",pady=pady,padx=padx)
    Letter_h = Button(master=master,bg=background,fg=foreground,text="H",command=lambda:Filler_func("H"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_h.grid(row=1,column=1,sticky="ew",pady=pady,padx=padx)

    Letter_i = Button(master=master,bg=background,fg=foreground,text="I",command=lambda:Filler_func("I"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_i.grid(row=1,column=2,sticky="ew",pady=pady,padx=padx)

    Letter_j = Button(master=master,bg=background,fg=foreground,text="J",command=lambda:Filler_func("J"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_j.grid(row=1,column=3,sticky="ew",pady=pady,padx=padx)

    Letter_k = Button(master=master,bg=background,fg=foreground,text="K",command=lambda:Filler_func("K"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_k.grid(row=1,column=4,sticky="ew",pady=pady,padx=padx)

    Letter_l = Button(master=master,bg=background,fg=foreground,text="L",command=lambda:Filler_func("L"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_l.grid(row=1,column=5,sticky="ew",pady=pady,padx=padx)


    Letter_m = Button(master=master,bg=background,fg=foreground,text="M",command=lambda:Filler_func("M"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_m.grid(row=2,column=0,sticky="ew",pady=pady,padx=padx)

    Letter_n = Button(master=master,bg=background,fg=foreground,text="N",command=lambda:Filler_func("N"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_n.grid(row=2,column=1,sticky="ew",pady=pady,padx=padx)

    Letter_o = Button(master=master,bg=background,fg=foreground,text="O",command=lambda:Filler_func("O"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_o.grid(row=2,column=2,sticky="ew",pady=pady,padx=padx)

    Letter_p = Button(master=master,bg=background,fg=foreground,text="P",command=lambda:Filler_func("P"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_p.grid(row=2,column=3,sticky="ew",pady=pady,padx=padx)

    Letter_q = Button(master=master,bg=background,fg=foreground,text="Q",command=lambda:Filler_func("Q"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_q.grid(row=2,column=4,sticky="ew",pady=pady,padx=padx)
    Letter_r = Button(master=master,bg=background,fg=foreground,text="R",command=lambda:Filler_func("R"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_r.grid(row=2,column=5,sticky="ew",pady=pady,padx=padx)

    Letter_s = Button(master=master,bg=background,fg=foreground,text="S",command=lambda:Filler_func("S"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_s.grid(row=3,column=0,sticky="ew",pady=pady,padx=padx)


    Letter_t = Button(master=master,bg=background,fg=foreground,text="T",command=lambda:Filler_func("T"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_t.grid(row=3,column=1,sticky="ew",pady=pady,padx=padx)

    Letter_u = Button(master=master,bg=background,fg=foreground,text="U",command=lambda:Filler_func("U"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_u.grid(row=3,column=2,sticky="ew",pady=pady,padx=padx)

    Letter_v = Button(master=master,bg=background,fg=foreground,text="V",command=lambda:Filler_func("V"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_v.grid(row=3,column=3,sticky="ew",pady=pady,padx=padx)

    Letter_w = Button(master=master,bg=background,fg=foreground,text="W",command=lambda:Filler_func("W"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_w.grid(row=3,column=4,sticky="ew",pady=pady,padx=padx)

    Letter_x = Button(master=master,bg=background,fg=foreground,text="X",command=lambda:Filler_func("X"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_x.grid(row=3,column=5,sticky="ew",pady=pady,padx=padx)

    Letter_y = Button(master=master,bg=background,fg=foreground,text="Y",command=lambda:Filler_func("Y"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_y.grid(row=4,column=2,sticky="ew",pady=pady,padx=padx)

    Letter_z = Button(master=master,bg=background,fg=foreground,text="Z",command=lambda:Filler_func("Z"),width=width,height=height,overrelief="groove",borderwidth=borderwidth,activebackground=activebg)
    Letter_z.grid(row=4,column=3,sticky="ew",pady=pady,padx=padx)
    window.bind("<KeyPress>", lambda event: Check_letter(event.keysym) if event.char.isalpha() else print(""))
    ButtonName_Letter = {"a":Letter_a ,
                        "b":Letter_b ,
                        "c":Letter_c ,
                        "d":Letter_d ,
                        "e":Letter_e ,
                        "f":Letter_f ,
                        "g":Letter_g ,
                        "h":Letter_h ,
                        "i":Letter_i ,
                        "j":Letter_j ,
                        "k":Letter_k ,
                        "l":Letter_l ,
                        "m":Letter_m ,
                        "n":Letter_n ,
                        "o":Letter_o ,
                        "p":Letter_p ,
                        "i":Letter_i ,
                        "q":Letter_q ,
                        "r":Letter_r ,
                        "s":Letter_s ,
                        "t":Letter_t ,
                        "u":Letter_u ,
                        "v":Letter_v ,
                        "w":Letter_w ,
                        "x":Letter_x ,
                        "y":Letter_y ,
                        "z":Letter_z
                        }
