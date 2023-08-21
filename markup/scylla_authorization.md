#scylla.yaml
    authorizer: CassandraAuthorizer

sudo systemctl restart scylla-server

#Create a role for the superuser which has all privileges

    CREATE ROLE <role-name> WITH SUPERUSER = true;
    CREATE ROLE DBA WITH SUPERUSER = true;

#Assign that role to yourself and grant login privileges

    CREATE ROLE <user> WITH PASSWORD = 'password' AND SUPERUSER = true AND LOGIN = true;

#For example (John is the DBA)

    CREATE ROLE john WITH PASSWORD = '39fksah!' AND LOGIN = true;
    GRANT DBA TO john;


#REVOKE ROLE
    REVOKE `role_name` FROM `role_name`
    REVOKE report_writer FROM alice;

#Configure the appropriate access privileges for clients using GRANT PERMISSION statements. For additional examples, consult the RBAC example.
#In this example, you are creating a user (db_user) who can access with password (password). You are also granting db_user with the role named client who has SELECT permissions on the ks.t1 table.

    CREATE ROLE db_user WITH PASSWORD = 'password' AND LOGIN = true;
    CREATE ROLE client;
    GRANT SELECT ON <keyspace1.table1> TO client;
    GRANT client TO db_user;


    GRANT MODIFY ON schedule.cust TO staff;


#Remove Cassandra Default Password and User

    DROP ROLE [ IF EXISTS ] 'old-username';
    DROP ROLE [ IF EXISTS ] 'cassandra';


#LIST ROLES
    LIST ROLES;
    LIST ROLES OF alice;
    LIST ROLES OF bob NORECURSIVE;



CREATE ROLE admin WITH PASSWORD = 'algo83111%%' AND SUPERUSER = true AND LOGIN = true;

CREATE ROLE lss WITH PASSWORD = 'algo83111%%' AND LOGIN = true;
CREATE ROLE app_user;
GRANT MODIFY ON s1.comments TO app_user;
GRANT app_user TO lss;




LIST ALL PERMISSIONS OF lss;