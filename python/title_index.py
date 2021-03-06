# -*- coding: utf-8 -*-

# Utility for building a title list for a standard history from Wikisource.
import codecs


FOOTER = """\n\n
Chinese text: This work was published before January 1, 1923, and is in the
public domain worldwide because the author died at least 100 years ago.
"""

def write_titles():
  dname = '../corpus/'
  prefix = 'songshi'
  with codecs.open('titles.txt', 'r', "utf-8") as title_file:
    with codecs.open('temp.txt', 'w', "utf-8") as f:
      i = 1
      j = 1
      k = 1
      l = 1
      for line in title_file:
        tokens = line.split()
        #print(tokens)
        fname = ''
        english = u'Volume %d' % i
        if len(tokens) > 2 and tokens[1].startswith(u'本紀'):
          english += ' Annals %d' % i
        elif len(tokens) > 2 and tokens[1].startswith(u'志'):
          english += ' Treatises %d' % j
          j += 1
        elif len(tokens) > 2 and u'列傳' in tokens[1]:
          english += ' Biographies %d' % k
          k += 1
        elif len(tokens) > 2 and tokens[1].startswith(u'表'):
          english += ' Tables %d' % l
          l += 1
        line = line.strip()
        title = u'\t%s %s: \n' % (line, english)
        if i < 10:
          fname = u'{0}/{1}00{2}.txt'.format(prefix, prefix, i)
          f.write(u'%s\t%s/%s00%d.html%s' % (fname, prefix, prefix, i, title))
        elif i < 100:
          fname = u'{0}/{1}0{2}.txt'.format(prefix, prefix, i)
          f.write(u'%s\t%s/%s0%d.html%s' % (fname, prefix, prefix, i, title))
        else:
          fname = u'{0}/{1}{2}.txt'.format(prefix, prefix, i)
          f.write(u'%s\t%s/%s%d.html%s' % (fname, prefix, prefix, i, title))
        path = '{0}{1}'.format(dname, fname)
        g = open(path, 'w')
        g.write(FOOTER)
        i += 1


def main():
  write_titles()


if __name__ == "__main__":
  main()
