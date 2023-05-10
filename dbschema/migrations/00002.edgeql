CREATE MIGRATION m1tfupk3yv4k4j3vuc2m3nbmle5e5turrkewm47remai6k74m2ilfq
    ONTO m1ac5tovfcyhwn75efsvgi5rmwyuhsorn5jvoewbcaamo5kxlv3yvq
{
  ALTER TYPE default::Book {
      ALTER PROPERTY release_tear {
          RENAME TO release_date;
      };
  };
};
