CREATE MIGRATION m1yom2uhbzghsvnwwgzzs4p5pyetpx6x3hqhg3adce5ibd4pc6tnra
    ONTO initial
{
  CREATE ABSTRACT TYPE default::BaseObject {
      CREATE REQUIRED PROPERTY created_at: std::datetime {
          SET readonly := true;
          CREATE REWRITE
              INSERT 
              USING (std::datetime_of_statement());
      };
      CREATE PROPERTY deleted_at: std::datetime;
      CREATE REQUIRED PROPERTY updated_at: std::datetime {
          SET readonly := true;
          CREATE REWRITE
              INSERT 
              USING (std::datetime_of_statement());
          CREATE REWRITE
              UPDATE 
              USING (std::datetime_of_statement());
      };
  };
  CREATE TYPE default::User EXTENDING default::BaseObject {
      CREATE REQUIRED PROPERTY email: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY password: std::str;
      CREATE REQUIRED PROPERTY username: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
};
