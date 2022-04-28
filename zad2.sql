select * from (
	select d,round(avg(t1),2) t1,round(avg(t2),2) t2,round(avg(t1)-avg(t2),2) dt,round(sum(M1),3) M1,round(sum(M2),3) M2, round(sum(M1)-sum(m2),3) dM, round(sum(Q),3) Q , 24.00 BHP  from (

		select format(dateRecord,'dd.MM.yy') d,record t1,NULL t2,0 M1, 0 M2, 0 Q from zap where recordNameId=11
		union all
		select format(dateRecord,'dd.MM.yy') d,NULL t1, record t2,0 M1, 0 M2,0 Q from zap where recordNameId=12 
		union all
		select format(dateRecord,'dd.MM.yy') d,NULL t1, NULL t2,recordDiff M1, 0 M2,0 Q from zap where recordNameId=2 
		union all
		select format(dateRecord,'dd.MM.yy') d,NULL t1, NULL t2,0 M1, recordDiff M2,0 Q from zap where recordNameId=3 
		union all
		select format(dateRecord,'dd.MM.yy') d,NULL t1, NULL t2,0 M1, 0 M2,recordDiff Q from zap where recordNameId=1 
	
	) as tbl
	group by d
union all
	select 'ИТОГО' d,round(avg(t1),2) t1,round(avg(t2),2) t2,round(avg(t1)-avg(t2),2) dt,round(sum(M1),3) M1,round(sum(M2),3) M2, round(sum(M1)-sum(m2),3) dM, round(sum(Q),3) Q , 24.00 BHP  from (
		select record t1,NULL t2,0 M1, 0 M2, 0 Q from zap where recordNameId=11
		union all
		select NULL t1, record t2,0 M1, 0 M2,0 Q from zap where recordNameId=12 
		union all
		select NULL t1, NULL t2,recordDiff M1, 0 M2,0 Q from zap where recordNameId=2 
		union all
		select NULL t1, NULL t2,0 M1, recordDiff M2,0 Q from zap where recordNameId=3 
		union all
		select NULL t1, NULL t2,0 M1, 0 M2,recordDiff Q from zap where recordNameId=1 
	) as tbl2
) as tbl3
ORDER BY IIF(d='ИТОГО','3000-01-01',CONVERT(DATE, d)) ASC
