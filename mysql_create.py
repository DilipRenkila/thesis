#!/usr/bin/python

import MySQLdb

db = MySQLdb.connect("localhost","root","1","thesis" )

cursor = db.cursor()


# Create table as per requirement
sql = """CREATE TABLE experiments (expid int(11) NOT NULL AUTO_INCREMENT,
         runid int(11) NOT NULL,
         keyid int(11) NOT NULL,
         delay_on_shaper  CHAR(20) NOT NULL,
         packets_sent int(11) NOT NULL,
         min_packet_length int(11) NOT NULL,
         max_packet_lenth int(11) NOT NULL,
         packet_distribution  CHAR(20) NOT NULL,
         sampling_interval int(11) NOT NULL,
         min_intergramegap int(11) NOT NULL,
         max_intergramegap int(11) NOT NULL,
         interframegap_distribution  CHAR(20) NOT NULL,
         destination  CHAR(20) NOT NULL,
         status int(11) NOT NULL,
         PRIMARY KEY(expid)) AUTO_INCREMENT=1 """

cursor.execute(sql)

# disconnect from server
db.close()

