CREATE MIGRATION m1ac5tovfcyhwn75efsvgi5rmwyuhsorn5jvoewbcaamo5kxlv3yvq
    ONTO initial
{
  CREATE FUTURE nonrecursive_access_policies;
  CREATE TYPE default::Book {
      CREATE REQUIRED PROPERTY author -> std::str;
      CREATE REQUIRED PROPERTY genre -> std::str;
      CREATE PROPERTY quantity -> std::int16;
      CREATE REQUIRED PROPERTY release_tear -> std::datetime;
      CREATE PROPERTY sales -> std::bigint;
      CREATE REQUIRED PROPERTY title -> std::str;
  };
};
