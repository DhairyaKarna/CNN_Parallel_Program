import pandas as pd
import matplotlib.pyplot as plt

# Load the existing results data from 'results.xlsx'
results_df = pd.read_excel('results.xlsx', sheet_name='Sheet1')

# Create a list to hold the speedup data
speedup_list = []

# Loop through each unique combination of Data Directory and Parallel Mode
for (data_dir, parallel_mode), group in results_df.groupby(['Data Directory', 'Parallel Mode']):
    # Get the time taken with 1 thread
    time_1_thread = group[group['Thread Count'] == 1]['Fastest Time'].values[0]
    for index, row in group.iterrows():
        # Calculate the speedup
        speedup = time_1_thread / row['Fastest Time']
        speedup_data = {
            'Data Directory': data_dir,
            'Parallel Mode': parallel_mode,
            'Thread Count': row['Thread Count'],
            'Speedup': speedup
        }
        speedup_list.append(speedup_data)

# Create a DataFrame from the speedup list
speedup_df = pd.DataFrame(speedup_list)

# Save the speedup data to a new sheet in the Excel file
with pd.ExcelWriter('results.xlsx', mode='a') as writer:
    speedup_df.to_excel(writer, sheet_name='Speedup', index=False)

# Function to generate and save the speedup graph for a given Parallel Mode
def generate_graph(parallel_mode):
    mode_df = speedup_df[speedup_df['Parallel Mode'] == parallel_mode]
    plt.figure(figsize=(10, 6))
    for data_dir, group in mode_df.groupby('Data Directory'):
        plt.plot(group['Thread Count'], group['Speedup'], label=data_dir)
    plt.title(f'Speedup Graph for {parallel_mode}')
    plt.xlabel('Number of Threads')
    plt.ylabel('Speedup')
    plt.legend()
    plt.grid(True)
    plt.savefig(f'speedup-{parallel_mode}.png')

# Generate the speedup graphs
generate_graph('parfiles')
generate_graph('parslices')