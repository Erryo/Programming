import  tkinter as tk
from tkinter import *
root = tk.Tk()
root.geometry("300x120")
root.config(bg="#203f5f")
root.title("Converter")
up_unit = tk.Variable(value="kg")
cur_unit = tk.Variable(value="kg")
low_unit = tk.Variable(value="kg")
main_frame = tk.Frame(master=root,bg="white")
from_frame = tk.Frame(master=main_frame,bg="blue")
to_frame = tk.Frame(master=main_frame,bg="red")

Convert_dict_mass = {("kg","g"):1000,
                ("kg","t"):1/1000,
                ("kg","mg"):1000000,
                ("kg","μg"):1e+6,
                ("kg","imperial ton"):1/1016,
                ("kg","Pound"):2.205,
                ("kg","Ounce"):35.274,

}
Convert_dict_volume ={
# Volume conversions
                ("liter", "ml"): 1000,
                ("liter", "m^3"): 0.001,
                ("liter", "gallon"): 0.264172,
                ("liter", "quart"): 1.056688,

}
Convert_dict_length ={
    # Length
    ("meter", "mm"): 1000,
    ("meter", "cm"): 100,
    ("meter", "km"): 0.001,
    ("meter", "mile"): 0.000621371,
    ("meter", "yard"): 1.09361,
    ("meter", "foot"): 3.28084,
    ("meter", "inch"): 39.3701,

}
Convert_dict_pressure ={

    # Pressure
    ("pascal", "kPa"): 0.001,
    ("pascal", "MPa"): 1e-6,
    ("pascal", "bar"): 1e-5,
    ("pascal", "atm"): 9.8692e-6,
    ("pascal", "mmHg"): 0.00750062,

}
Convert_dict_speed={

    # Speed
    ("m/s", "km/h"): 3.6,
    ("m/s", "mph"): 2.23694,

}
Convert_dict_Angle ={
    # Angle
    ("degree", "radian"): 0.0174533,
    ("degree", "gradian"): 1.11111,

}
Convert_dict_time={

    # Time
    ("second", "minute"): 1 / 60,
    ("second", "hour"): 1 / 3600,
    ("minute", "hour"): 1 / 60,

}
Convert_dict_temperature={
    # Temperature
    ("celsius", "fahrenheit"): (9 / 5, 32),
    ("celsius", "kelvin"): 1,

    ("fahrenheit", "celsius"): (5 / 9, -32 * (5 / 9)),
    ("fahrenheit", "kelvin"): (5 / 9, 273.15),

    ("kelvin", "celsius"): 1,
    ("kelvin", "fahrenheit"): (9 / 5, -32 * (5 / 9)),}
def Convert():
    measure = Measure_var.get()
    value = float(from_quant_entry.get())
    from_unit = str(from_unit_var.get())
    to_unit = to_unit_var.get()
    print(value)
    try:
        result = eval("value * float(Convert_dict_"+measure+".get((from_unit,to_unit)))")
        print(result)
        result_var.set(result)
    except ValueError:
        pass
def Change_units(event):
    global to_unit_drpdwn,from_unit_drpdwn
    from_unit_drpdwn.pack_forget()
    from_unit_drpdwn = eval("OptionMenu(from_frame,from_unit_var,*Units_"+Measure_var.get()+")")
    to_unit_drpdwn.pack_forget()
    to_unit_drpdwn = to_unit_drpdwn = eval("OptionMenu(to_frame,to_unit_var,*Units_"+Measure_var.get()+")")
    to_unit_drpdwn.pack()
    from_unit_drpdwn.pack()
Units_mass =["t","kg","g","mg","μg","imperial ton","Pound","Ounce"]
Units_volume =["liter","ml","m^3","gallon","quart","pint","cup","Fluid Ounce"]
Units_length=["m","km","mm","cm","mile","yard","foot","inch"]
Units_pressure=["pascal","kPa","MPa","bar","atm","mmHg"]
Units_speed=["m/s","km/h","mph"]
Units_Angle=["degree","radian","gradian"]
Units_time=["second","minute","hour"]
Units_temperature=["celsius","fahrenheit","kelvin","mg","μg","imperial ton","Pound","Ounce"]
Units = Units_mass
Measures =["mass","volume","length","pressure","speed","Angle","time","temperature"]
from_unit_var = StringVar(master=root,value="kg")
Measure_var = StringVar(master=root,value="mass")
to_unit_var = StringVar(master=root,value="kg")
result_var = Variable(master=root,value="")

from_quant_entry = Entry(master=from_frame)
from_unit_drpdwn = OptionMenu(from_frame,from_unit_var,*Units_mass)


To_quant_Label = Label(master=to_frame,textvariable=result_var)
to_unit_drpdwn = OptionMenu(to_frame,to_unit_var,*Units_mass)

Conver_button= Button(master= root,text="convert",command=Convert)
Measure_drpdwn =  OptionMenu(root,Measure_var,*Measures,command=Change_units)
Measure_drpdwn.pack(side="top")
from_quant_entry.pack(side="top")
from_unit_drpdwn.pack(side="top")
To_quant_Label.pack(side="top")
to_unit_drpdwn.pack(side="top")
from_frame.pack(side="left")
to_frame.pack(side="left")
Conver_button.pack(side="bottom")
main_frame.pack()

root.mainloop()