CREATE TABLE [dbo].[zap](
	[id] [bigint] NOT NULL,
	[counterId] [bigint] NOT NULL,
	[recordNameId] [bigint] NOT NULL,
	[dateRecord] [datetime] NOT NULL,
	[record] [float] NOT NULL,
	[recordDiff] [float] NOT NULL,
	[period] [nchar](6) NOT NULL
) ON [PRIMARY]
GO

CREATE TABLE [dbo].[spr](
	[id] [bigint] NOT NULL,
	[counterId] [bigint] NOT NULL,
	[recordName] [nchar](64) NOT NULL,
	[isCalculateDiff] [tinyint] NOT NULL,
	[recordNameRus] [nchar](512) NOT NULL
) ON [PRIMARY]
GO

--truncate table zap
BULK INSERT zap
FROM 'C:\1\zap.csv'
WITH
(
FirstRow = 2,
FIELDTERMINATOR = ';',
ROWTERMINATOR = '\n'
);
select* from zap

--truncate table spr
BULK INSERT spr
FROM 'C:\1\spr.csv'
WITH
(
FirstRow = 2,
FIELDTERMINATOR = ';',
ROWTERMINATOR = '\n',
CODEPAGE = '1251'
);
select* from spr

/*
insert into zap
(id,counterId,recordNameId,dateRecord,record,recordDiff,period)
values (16223,1069,1,'2021-01-01 00:00:00.000',28547.2376259416,36.4249167889357,'202101')
*/
