

import pandas as pd
import matplotlib.pyplot as plt


if __name__ == "__main__":
    data = pd.read_csv('testlogfile', sep=" ", header=None)

    selectedDf = data[[1,2]]
    
    voltages =  selectedDf[selectedDf[2].str.endswith('V')]
    currents =  selectedDf[selectedDf[2].str.endswith('A')]
    currents[2] = currents[2].str.rstrip('A')
    currents[2] = pd.to_numeric(currents[2])
    
    print(voltages)
    print(currents)

    plt.figure()
    currents.plot()
    plt.show()
       

    