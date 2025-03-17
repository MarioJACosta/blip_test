# **Repository Activity Ranking Algorithm**

## **Overview**
This algorithm calculates an **activity score** for repositories based on commit history from a csv file. The score is computed using the number of **files changed**, **lines added**, and **lines deleted**. The algorithm then ranks the repositories based on their activity scores and outputs the **top 10 most active repositories**.

## **Algorithm Description**

1. **Read the CSV File**  
   - The CSV file is processed **line by line** using Go's `csv.Reader`.

2. **Calculate Activity Scores**  
   - For each repository, the activity score is calculated as:
     
     ```
     ActivityScore = filesChanged + additions + deletions
     ```
   - These scores are stored in a `map[string]int` where:
     - Key = Repository name
     - Value = Activity score

3. **Sort Repositories by Score**  
   - The repository scores are converted into a slice of `Repository` structs.
   - The list is then sorted in descending order based on activity scores using Go’s `sort.Slice()` function.
   
4. **Output the Top 10 Repositories**  
   - The top 10 repositories are printed to the console.

## **Implementation Details**
### **Activity Score Calculation**
The algorithm extracts the following fields from the CSV:
- **Repository Name** (Column 3)
- **Files Changed** (Column 4)
- **Additions** (Column 5)
- **Deletions** (Column 6)

The final score is computed as:
```go
repoScores[repoName] += filesChanged + additions + deletions
```

### **Sorting Repositories**
After calculating the activity scores, repositories are sorted using:
```go
sort.Slice(repositories, func(i, j int) bool {
    return repositories[i].ActivityScore > repositories[j].ActivityScore
})
```

## **Usage Instructions**
### **Prerequisites**
- Go (Golang) installed on your system.
- A CSV file (`commits.csv`) in the same directory as the script, containing commit data with the following columns:
  1. Timestamp
  2. User
  3. Repository Name
  4. Files Changed
  5. Additions
  6. Deletions

### **Running the Program**
To execute the program, use the following command:
```sh
go run main.go
```
This will display the **top 10 most active repositories**.