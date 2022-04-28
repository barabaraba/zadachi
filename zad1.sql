/*CREATE OR ALTER PROCEDURE report1
	@sprid int,
	@dA datetime,
	@dB datetime,
	@dFormat varchar(10)
AS
BEGIN		

	SELECT spr.recordNameRus,format(zap.dateRecord,@dFormat),
	IIF(@sprid in (11,12),AVG(IIF(spr.isCalculateDiff=1,zap.recordDiff,zap.record)),SUM(IIF(spr.isCalculateDiff=1,zap.recordDiff,zap.record))) val
	from zap,spr 
	where zap.recordNameId = spr.id
	and recordNameId=@sprid
	and dateRecord between @dA and @dB
	group by spr.recordNameRus,format(zap.dateRecord,@dFormat) 
	
END
GO*/

EXEC report1 11,'2021-01-01','2022-01-01','yyyy-MM'
