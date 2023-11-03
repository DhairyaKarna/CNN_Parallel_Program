import subprocess
import pandas as pd

# Define the parameters
data_dirs = ['small', 'mixture', 'big']
parallel_modes = ['parfiles', 'parslices']
thread_counts = [1, 2, 4, 6, 8, 12]

results_list = []

# Loop through each combination of parameters
for data_dir in data_dirs:
    for parallel_mode in parallel_modes:
        for thread_count in thread_counts:
            # Create a list to hold the time results for each run
            run_times = []
            for i in range(5):  # Run each configuration 5 times
                # Build the command string
                cmd = f'go run ../editor/editor.go {data_dir} {parallel_mode} {thread_count}'
                # Run the command and capture the output
                result = subprocess.run(cmd, shell=True, text=True, stdout=subprocess.PIPE)
                # Extract the time printed from editor.go as the output
                time_taken = float(result.stdout.strip())
                run_times.append(time_taken)
            # Find the fastest time and save the result
            fastest_time = min(run_times)
            result_data = {
                'Data Directory': data_dir,
                'Parallel Mode': parallel_mode,
                'Thread Count': thread_count,
                'Fastest Time': fastest_time
            }
            results_list.append(result_data)

# Create a DataFrame to hold the results
columns = ['Data Directory', 'Parallel Mode', 'Thread Count', 'Fastest Time']
results_df = pd.DataFrame(results_list, columns=columns)

# Save the results to a spreadsheet
results_df.to_excel('results.xlsx', index=False)
