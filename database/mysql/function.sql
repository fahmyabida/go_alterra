CREATE FUNCTION DELETE_HISTORY_TRANSACTION (transactionId int)
RETURNS text
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = transactionId;
    RETURN 'success';
END;

SET GLOBAL log_bin_trust_function_creators = 1;

SELECT CREATE_PRODUCT('drink', 1);