CREATE MIGRATION m1zn24olkul4ovuom4kddv6k2zkdzgygry32jmrcgzj7ueyk4o2nqa
    ONTO m1tfupk3yv4k4j3vuc2m3nbmle5e5turrkewm47remai6k74m2ilfq
{
  ALTER TYPE default::Book {
      ALTER PROPERTY quantity {
          SET REQUIRED USING (1000);
      };
  };
};
