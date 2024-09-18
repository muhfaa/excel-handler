# API Documentation: Excel File Upload Endpoint

This documentation describes how to use the API to upload and process Excel (.xlsx) files. The API supports files with varying structures, including different numbers of columns, rows, and data types.

## How to Run This Service

1. **Prerequisites:**
   - Install [Go](https://golang.org/dl/).
   - Install the required Go modules with `go mod tidy`.

2. **Setup:**
   - Clone or download the service code.
   - Navigate to the project directory.

3. **Start the Service:**
   - Run `go run main.go` to start the service.
   - The service listens on port 8080 by default.

4. **Testing:**
   - Use Postman or cURL to send a POST request with an Excel file to `http://localhost:8080/upload`.

## Endpoint Overview

- **URL**: `/upload`
- **Method**: POST
- **Content-Type**: multipart/form-data
- **Description**: Upload an Excel file for extraction and processing.

## Request Parameters

| Parameter | Type | Description                              |
|-----------|------|------------------------------------------|
| `file`    | file | The Excel file to upload (.xlsx format only). |

## Supported File Formats

- **Excel**: Only .xlsx files (Excel 2007 or later) are supported.

## Request Example

### Uploading an Excel file (.xlsx):

**Postman:**
1. Set the method to POST.
2. URL: `http://localhost:8080/upload`.
3. In the "Body" tab, select "form-data" and set the key to `file`. Upload an `.xlsx` file.

**cURL:**
```bash
curl -X POST http://localhost:8080/upload -F "file=@your_excel_file.xlsx"
```

## Response

The API will return a JSON response with the extracted data from the uploaded Excel file.

### Success Response

**Status**:  **200 OK**
**Response Example**:
```json
{
    "message": "Excel file uploaded and processed successfully",
    "data": [
        ["Header1", "Header2", "Header3"],
        ["Row1Col1", "Row1Col2", "Row1Col3"],
        ["Row2Col1", "Row2Col2", "Row2Col3"]
    ]
}
```

### Success Message
- **Message:** A success message indicating the Excel file was uploaded and processed.
- **Data:** A 2D array containing the extracted data from the Excel file. Each row of the array corresponds to a row from the file, with individual cell values in string format.

### Error Responses

| Error Type                      | Status                  | Response Example                                                    |
|---------------------------------|-------------------------|---------------------------------------------------------------------|
| Invalid File Type (e.g., unsupported format) | 400 Bad Request        | `json { "error": "Unsupported file format. Please upload an Excel (.xlsx) file." }` |
| Missing File                    | 400 Bad Request        | `json { "error": "File is required" }`                              |
| File Processing Failure         | 500 Internal Server Error | `json { "error": "Failed to process file" }`                |

### Handling Varying Excel File Structures

| Feature                  | Description                                                                                                               |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------|
| Unknown Column Counts    | The API dynamically processes each row based on the number of columns it contains. It doesn’t assume a fixed column count, allowing it to adapt to Excel files where some rows have more or fewer columns. |
| Mixed Data Types         | All data is extracted as a string for consistency. For example, whether a cell contains a number, date, or text, it will be processed as a string. You can further process this data on the client side if necessary. |
| Multiple Sheets          | The API processes data from all sheets in the Excel file. Data from each sheet is extracted and returned in the response, preserving the structure of each sheet. |
