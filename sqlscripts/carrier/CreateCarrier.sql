-- Create a new stored procedure called 'CreateCarrier' in schema 'dbo'
-- Drop the stored procedure if it already exists
IF EXISTS (
SELECT *
FROM INFORMATION_SCHEMA.ROUTINES
WHERE SPECIFIC_SCHEMA = N'dbo'
    AND SPECIFIC_NAME = N'CreateCarrier'
    AND ROUTINE_TYPE = N'PROCEDURE'
)
DROP PROCEDURE dbo.CreateCarrier
GO
-- Create the stored procedure in the specified schema
CREATE PROCEDURE dbo.CreateCarrier
    @CarrierCode NVARCHAR(50),
    @CarrierName NVARCHAR(50),
    @CarrierAddress NVARCHAR(50),
    @CarrierEmail NVARCHAR(50),
    @CarrierPhone NVARCHAR(50),
    @LastInsertedId INT OUTPUT
AS
BEGIN

    INSERT Carrier
        (
        CarrierCode,
        CarrierName,
        CarrierAddress,
        CarrierEmail,
        CarrierPhone
        )
    VALUES
        (
            @CarrierCode,
            @CarrierName,
            @CarrierAddress,
            @CarrierEmail,
            @CarrierPhone
    )
    SET @LastInsertedId = (SELECT CONVERT(INT, SCOPE_IDENTITY()))
END
GO