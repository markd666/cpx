

import pandas as pd
import matplotlib.pyplot as plt


if __name__ == "__main__":
    data = pd.read_csv('testlogfile', sep=" ", header=None)

    selectedDf = data[[1,2]]
    
    voltages =  selectedDf[selectedDf[2].str.endswith('V')]
    voltages[2] = voltages[2].str.rstrip('V')
    voltages[2] = pd.to_numeric(voltages[2])

    currents =  selectedDf[selectedDf[2].str.endswith('A')]
    currents[2] = currents[2].str.rstrip('A')
    currents[2] = pd.to_numeric(currents[2])
    
    print(voltages.loc[:,2])
    print(currents[2])
  
  
    ax = voltages.plot()
    ax.set_ylim(0, 50)
    ax.legend(["Voltage"])

    ay = currents.plot()
    ay.legend(["Current"])
    plt.show()
       


    