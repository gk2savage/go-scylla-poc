CQL Commands:

cqlsh> 

CREATE KEYSPACE avdb WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','datacenter1' : 1};

use avdb;


CREATE TABLE product_data (
   pid text,
   pname text,
   os text,
   product_location text,
   PRIMARY KEY((pid, pname)));

insert into product_data ("pid","pname","os","product_location") VALUES ('P101AN','zsecurity','Unix', 'http://www.nsdsec.com/zsecurity');

insert into product_data ("pid","pname","os","product_location") VALUES ('A212PN','asecure','Windows', 'http://www.secureav.com/asecure');

insert into product_data ("pid","pname","os","product_location") VALUES ('CNQ12','asecure+','Windows', 'http://www.secureav.com/asecurepro');



SELECT * FROM product_data;

DROP KEYSPACE product;

DROP TABLE product_data;


///
CREATE KEYSPACE product
WITH replication = {'class': 'NetworkTopologyStrategy', 'DC1' : 1, 'DC2' : 3}
AND durable_writes = true;


ALTER KEYSPACE product
 WITH replication = { 'class' : 'NetworkTopologyStrategy', 'datacenter1' : 1};
///


