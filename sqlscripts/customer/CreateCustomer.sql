-- Create a new stored procedure called 'CreateCustomer' in schema 'dbo'
-- Drop the stored procedure if it already exists
IF EXISTS (
SELECT *
FROM INFORMATION_SCHEMA.ROUTINES
WHERE SPECIFIC_SCHEMA = N'dbo'
    AND SPECIFIC_NAME = N'CreateCustomer'
    AND ROUTINE_TYPE = N'PROCEDURE'
)
DROP PROCEDURE dbo.CreateCustomer
GO
-- Create the stored procedure in the specified schema
CREATE PROCEDURE dbo.CreateCustomer
    @CustomerCode NVARCHAR(50),
    @CustomerName NVARCHAR(50),
    @CustomerAddress NVARCHAR(50),
    @CustomerEmail NVARCHAR(50),
    @CustomerPhone NVARCHAR(50),
    @LastInsertedId INT OUTPUT
AS
BEGIN

    INSERT Customer
        (
        CustomerCode,
        CustomerName,
        CustomerAddress,
        CustomerEmail,
        CustomerPhone
        )
    VALUES
        (
            @CustomerCode,
            @CustomerName,
            @CustomerAddress,
            @CustomerEmail,
            @CustomerPhone
    )
    SET @LastInsertedId = (SELECT CONVERT(INT, SCOPE_IDENTITY()))
END
GO